package middleware

import (
	"log"
	"net/http"
)

func Logs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method : [%s] Path: [%s] RemoteAddr: [%v]\n", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
