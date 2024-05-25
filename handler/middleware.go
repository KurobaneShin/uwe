package handler

import "net/http"

func WithAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("comming from middleware")
		next.ServeHTTP(w, r)
	})
}
