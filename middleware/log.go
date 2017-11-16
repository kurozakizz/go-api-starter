package middleware

import (
	"log"
	"net/http"
)

// WriteRequestLog will echo request log to console
func WriteRequestLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)
		next.ServeHTTP(w, r)
	}
}
