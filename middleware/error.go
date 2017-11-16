package middleware

import (
	"log"
	"net/http"
)

// TraceError will echo error to console
func TraceError(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("error", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
