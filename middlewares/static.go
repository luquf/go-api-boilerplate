package middlewares

import "net/http"

func StaticBlockMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" ||
			r.URL.Path == "images/" ||
			r.URL.Path == "js/" ||
			r.URL.Path == "assets/" ||
			r.URL.Path == "css/" {
			//NotFound404handler(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
