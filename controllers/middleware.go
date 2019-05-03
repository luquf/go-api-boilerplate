package controllers

import (
	"log"
	"net/http"
)

func StaticBlockMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" ||
			r.URL.Path == "images/" ||
			r.URL.Path == "js/" ||
			r.URL.Path == "assets/" ||
			r.URL.Path == "css/" {
			NotFound404handler(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// LogMiddleware handles request logging
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// LogoutMiddleware handles disauthentication
func LogoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:   "TODO :token_name",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)
		next.ServeHTTP(w, r)
	})
}

func AdminRedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := getAdminFsromCookie(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		RedirectHTTP(w, r, "/", 302)
		return

	})
}

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := getAdminFsromCookie(r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next.ServeHTTP(w, r)
	})
}
