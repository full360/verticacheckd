package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gitlab.full360.com/full360/verticacheckd"
)

func main() {
	addr := flag.String("address", "127.0.0.1", "HTTP address")
	port := flag.Int("port", 3000, "HTTP server listening port")
	name := flag.String("service-name", "verticacheckd", "Service name")
	timeOut := flag.Duration("timeout", 5*time.Second, "HTTP Read and Write timeout")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()

	// log setup
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime)

	hostAddr, err := verticacheckd.ExternalIP()
	if err != nil {
		logger.Fatal(err)
	}

	svc := verticacheckd.NewService(
		hostAddr,
		"/opt/vertica/bin/admintools",
		[]string{"-t", "view_cluster", "-x"},
	)

	r := mux.NewRouter()
	s := r.PathPrefix(fmt.Sprintf("/%s", *name)).Subrouter()

	s.Handle(
		"/state",
		verticacheckd.StateHandler(svc),
	).Methods("GET")

	s.Handle(
		"/dbs/{name}/state",
		verticacheckd.DBStateHandler(svc),
	).Methods("GET")

	srv := http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", *addr, *port),
		WriteTimeout: *timeOut,
		ReadTimeout:  *timeOut,
	}

	logger.Printf("listening on address: %s, port: %d\n", *addr, *port)
	logger.Fatal(srv.ListenAndServe())
}
