package handler

import "net/http"

func WithAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		if len(apiKey) == 0 {
			writeUnAuthorized(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeUnAuthorized(w http.ResponseWriter) error {
	return writeJSON(w, http.StatusUnauthorized, UnAuthorized())
}
