package main

import (
	"go-api-boilerplate/config"
	"go-api-boilerplate/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	// var err error
	// var f *os.File
	// f, err := os.OpenFile("/var/log/api_name.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()

	// log.SetOutput(f)

	var conf = config.Configurations

	router := mux.NewRouter()

	router.StrictSlash(false)
	// Serve static assets
	// fileServer := http.FileServer(http.Dir("public"))
	// http.Handle("/static/", http.StripPrefix("/public/", controllers.StaticBlockMiddleware(fileServer)))

	// TODO: DATABASE CONNECTION

	// Import all the api endpoints
	for _, route := range routes.Scheme {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	// 404 handler
	// router.NotFoundHandler = http.HandlerFunc(controllers.NotFound404handler)

	http.Handle("/", router)
	log.Println("Listening on 127.0.0.1:" + conf.AppPort)
	log.Fatal(http.ListenAndServe(conf.AppHost, nil))
}
