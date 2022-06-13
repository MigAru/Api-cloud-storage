package middlewares

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %q\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}