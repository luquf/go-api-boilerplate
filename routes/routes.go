package routes

import (
	"go-api-boilerplate/controllers"
	"go-api-boilerplate/middlewares"
	"net/http"
)

type route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type routes []route

var Scheme = routes{
	route{
		"Hello, World !",
		"GET",
		"/",
		middlewares.
			LogMiddleware(&controllers.HelloWorld{}),
	},
}
