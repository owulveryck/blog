package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
	"net"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func NodeDelete(w http.ResponseWriter, r *http.Request) {
}
func NodeDisplay(w http.ResponseWriter, r *http.Request) {
}
func NodeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	request := PublishNodeRequest(todoId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(request); err != nil {
		panic(err)
	}
	return

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -s -X POST -H 'Content-Type:application/json' -H 'Accept:application/json' -d '{"kind":"linux","size":"S","disksize":20,"leasedays":1,"environment_type":"dev","centrify_zone":"zd_fr_cld_01","description":"my_description","usergroup":"clddevsudo01","service_account":"poctst01","app_trigram":"TRG","region":"eu-fr-paris","availability_zone":"eu-fr-paris-1","subnet":"192.160.64.0/21","notifymail":"user1@example.com,user2@example.com"}' -k https://localhost:8080/v2/nodes -u

*/
func NodeCreate(w http.ResponseWriter, r *http.Request) {
	var nodeRequest NodeRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &nodeRequest); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	conn, err := net.Dial("tcp", "127.0.0.1:18800")
	if err != nil {
		fmt.Println("fail to connect to server.")
		return

	}
	client := rpc.NewSession(conn, true)
	t, err := client.Send("NodeCreate", nodeRequest.Kind, nodeRequest.Size, nodeRequest.Disksize, nodeRequest.Leasedays, nodeRequest.EnvironmentType, nodeRequest.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return

	}
	fmt.Println(t)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
