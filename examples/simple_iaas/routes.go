package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"NodeCreate",
		"POST",
		"/v1/nodes",
		NodeCreate,
	},
	Route{
		"NodeDisplay",
		"GET",
		"/v1/nodes",
		NodeDisplay,
	},
	Route{
		"NodeShow",
		"GET",
		"/v1/nodes/{todoId}",
		NodeShow,
	},
	Route{
		"NodeDelete",
		"DELETE",
		"/v1/nodes/{todoId}",
		NodeDelete,
	},
}
