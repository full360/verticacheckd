package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("application is healthy"))
}

func Status(w http.ResponseWriter, t *http.Request) {
	w.Write([]byte("vertica status"))
}

func main() {
	addr := flag.String("address", "127.0.0.1", "HTTP address")
	port := flag.Int("port", 3000, "HTTP server listening port")
	name := flag.String("service-name", "vertica-check", "Service name")

	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("/%s/health", *name), Chain(Health, logging()))
	mux.HandleFunc(fmt.Sprintf("/%s/status", *name), Chain(Status, logging()))

	srv := http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf("%s:%d", *addr, *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
