package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes = Routes{
	{"Index", "GET", "/", Index},
	{"TodoIndex", "GET", "/todos", TodoIndex},
	{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
	{"TodoCreate", "POST", "/todos", TodoCreate},
}
