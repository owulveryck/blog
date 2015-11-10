+++
date = "2015-11-10T08:56:36+01:00"
draft = true
title = "a IaaS-like RESTfull API based on microservices"

+++

# Absract

Recently, I've been looking at the principles of a middleware layer and especially on how [API would glue a system](http://insertpulereference).

I've also seen this excellent video made by [Mat Ryer](http://reference) about how to code an API in GO and why go would be the perfect language to code such a portal.

The problem I'm facing is that in the organization I'm working for, the developments are heterogenous and therefore you can find *ruby* teams as well as *python* team and myself as a *go* team (That will change in the future anyway)
The point is that I would like my middleware to serve as an entry point to the services provided by the departement.

We (as a department) would then be able to present the interface via, for example, a [swagger](http://swagger.io) like interface, take care of the API and do whatever RPC to any submodule.

# An example: a IAAS like interface

Let's consider a node compute lifecycle.
What I want to do is :

* create a node
* update a node (maybe)
* delete a node
* get the status of the node

## The backend

The backend is whatever service able to create a node, suchs as openstack, vmware vcac, juju, or whatever. 
Thoses services usually provide RESTfull API.

What I've seen in my experience, is that usually, the API are given with a library in whatever modern language. 
This aim to simplify the developpement of the clients.
Sometimes this library may also be developped by an internal team that will take care of the maintenance.

## The library

In my example, we will consider that the library is a simple _gem_ file developped in ruby. 
Therefore, our service will be a simple server that will get RPC calls, call the good method in the _gemfile_ 
and that will, _in fine_ transfer it to the backend.

## The RestFull API.

I will use the example described [here](http://blogpost) as a basis for my work.

## The glue: MSGPACK-RPC

There are severeal method for RPC-ing different languages. Ages ago, there was xml-rpc; then there has been json-rpc; 
I will use msgpack-rpc which is a binary, json base codec.
The communication between the Go Server and the ruby client will be donc over TCP via HTTP for example.

Later on, outside of the scope of this post, I may use ZMQ (as I have already blogged about 0MQ communication between thoses languages).

# The implementation

I will describe here the node creation via a POST method, and consider that the other method could be implemented in a similar way.

## The signature of the node creation

Here is the expected signature for creating a compute element:

```json
{
    "kind":"linux",
    "size":"S",
    "disksize":20,
    "leasedays":1,
    "environment_type":"dev",
    "description":"my_description",
}
```

The corresponding GO structure is:

```go
type NodeRequest struct {
    Kind string `json:"kind"` // Node kind (eg linux)
    Size string `json:"size"` // size
    Disksize         int    `json:"disksize"`
    Leasedays        int    `json:"leasedays"`
    EnvironmentType  string `json:"environment_type"`
    Description      string `json:"description"`
}
```

## The route

The Middleware is using the [gorilla mux package](http://gorilla.mux.io). 
According the description, I will add an entry in the routes array (into the _routes.go_ file):

```go
Route{
    "NodeCreate",
    "POST",
    "/v1/nodes",
    NodeCreate,
},
```

*Note* : I am using a prefix `/v1` for my API, for exploitation purpose.

I will then create the corresponding handler in the file with this signature

```go
func NodeCreate(w http.ResponseWriter, r *http.Request){
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
}
```

That's in this fuction that will be implemented RPC (client part). To keep it simple at the begining, 
I will instanciate a TCP connection on every call.
Don't throw things at me, that will be changed later following the advice of Mat Ryer.

## The implementation of the handler

### The RPC part

To use _msgpack_ I need to import the go implemtation `github.com/msgpack-rpc/msgpack-rpc-go/rpc`.
This library will take care of the encoding/decoding of the messages.

Let's dial the RPC server and call the `NodeCreate` method with, as argument, the information we had from the JSON input

```go
    conn, err := net.Dial("tcp", "127.0.0.1:18800")
    if err != nil {
        fmt.Println("fail to connect to server.")
        return
    }
    client := rpc.NewSession(conn, true)
    retval, err := client.Send("NodeCreate", 2, 3)
    if err != nil {
    fmt.Println(err)
        return
    }
    fmt.Println(rpc.CoerceInt(retval))

```
