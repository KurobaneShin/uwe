package handler

import (
	"net/http"

	"uwe/db"
)

func WithAuthentication(db db.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("Authorization")
			if len(apiKey) == 0 {
				writeUnAuthorized(w)
				return
			}

			_, err := db.GetAccountByAPIKey(apiKey)
			if err != nil {
				writeUnAuthorized(w)
				return

			}
			next.ServeHTTP(w, r)
		})
	}
}

func writeUnAuthorized(w http.ResponseWriter) error {
	return writeJSON(w, http.StatusUnauthorized, UnAuthorized())
}
