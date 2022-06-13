package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log middleware")
		log.Printf("[%s] %q\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}