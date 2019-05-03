package controllers

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, templatePath string, data interface{}) {
	t, _ := template.ParseFiles(templatePath)
	t.Execute(w, data)
}

// RedirectHTTP redirects
func RedirectHTTP(w http.ResponseWriter, r *http.Request, dest string, code int) {
	http.Redirect(w, r, dest, code)
}
