package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"gitlab.full360.com/full360-south/verticacheckd"
)

func main() {
	addr := flag.String("address", "127.0.0.1", "HTTP address")
	port := flag.Int("port", 3000, "HTTP server listening port")
	name := flag.String("service-name", "verticacheckd", "Service name")
	timeOut := flag.Duration("timeouts", 5*time.Second, "HTTP Read and Write timeout")

	flag.Parse()

	hostAddr, err := verticacheckd.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}

	svc := verticacheckd.NewService(hostAddr, "admintools", []string{"-t", "view_cluster", "-x"})

	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("/%s/state", *name), verticacheckd.StateHandler(svc))

	srv := http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf("%s:%d", *addr, *port),
		WriteTimeout: *timeOut,
		ReadTimeout:  *timeOut,
	}

	log.Fatal(srv.ListenAndServe())
}
