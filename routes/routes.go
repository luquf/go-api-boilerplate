package routes

import (
	"net/http"
)

type route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type routes []route

var scheme = routes{
	// route{
	// 	"admin login page get",
	// 	"GET",
	// 	"/login",
	// 	controllers.
	// 		LogMiddleware(controllers.
	// 			AdminRedirectMiddleware(&controllers.AdminLogin{})),
	// }

}
