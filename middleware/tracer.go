package middleware

import (
	"log"
	"net/http"
)

// WriteTracer will echo tracer to console
func WriteTracer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("trace", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
