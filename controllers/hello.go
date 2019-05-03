package controllers

import "net/http"

// HelloWolrd: Boilerplate controller example
type HelloWorld struct{}

func (hw *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World !"))
	return
}
