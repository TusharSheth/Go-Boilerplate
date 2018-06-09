package main

import (
	"net/http"
)

// Route : route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
}
