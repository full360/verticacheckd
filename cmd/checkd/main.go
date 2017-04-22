package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	vcd "gitlab.full360.com/full360-south/verticacheckd"
)

func main() {
	addr := flag.String("address", "127.0.0.1", "HTTP address")
	port := flag.Int("port", 3000, "HTTP server listening port")
	name := flag.String("service-name", "verticacheckd", "Service name")
	timeOut := flag.Duration("timeouts", 5*time.Second, "HTTP Read and Write timeout")

	flag.Parse()

	hostAddr, err := vcd.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}

	cmd := "cat"
	args := []string{"fixture/cmd_output.txt"}

	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("/%s/health", *name), vcd.Check(hostAddr, cmd, args...))

	srv := http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf("%s:%d", *addr, *port),
		WriteTimeout: *timeOut,
		ReadTimeout:  *timeOut,
	}

	log.Fatal(srv.ListenAndServe())
}
