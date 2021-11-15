package middleware

import (
	"log"
	"net/http"
	"time"
)

func HelloWorldRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func TimedRoute(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	}
}