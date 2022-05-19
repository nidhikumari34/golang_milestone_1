package app

import (
	"log"
	"net/http"
	"time"
)

//logging execution time
func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime = time.Now()
		log.Printf("Execution start: %s", startTime.String())
		next.ServeHTTP(w, r)
		log.Printf("Execution end: %s", endTime.String())
		log.Printf("Total execution time: %s", execTime.String())
	}
}

//token verification
func AuthTokenMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header

		if header.Get("Auth-Token") == "" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("AuthToken missing"))
			return
		}
		h.ServeHTTP(w, r)
	}
}
