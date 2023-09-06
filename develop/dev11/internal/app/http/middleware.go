package http

import (
	"log"
	"net/http"
	"time"
)

func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s %s", r.Method, r.RemoteAddr, r.URL.Path, duration)
	}
}
