package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call next handler (Controller)

		next.ServeHTTP(w, r)
		// After response
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
