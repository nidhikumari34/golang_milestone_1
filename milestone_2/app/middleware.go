package app

import (
	"log"
	"net/http"
	"time"
)

//middleware for logging execution time
func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		StartTime = time.Now()
		log.Printf("Execution start: %s", StartTime.String())
		next.ServeHTTP(w, r)
		log.Printf("Execution end: %s", EndTime.String())
		log.Printf("Total execution time: %s", ExecTime.String())
	}
}

//middleware for token fetching from header
func AuthTokenMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(("X-Auth-Token"))

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("X-Auth-Token missing"))
			return
		}

		if !ValidToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid X-Auth-Token"))
			return
		}
		h.ServeHTTP(w, r)
	}
}
