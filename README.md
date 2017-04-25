# Vertica State

Vertica State is a web service that returns the state (up/down) of the Vertica
node it's running on.

## Detail

It will receive a request from an AWS ELB and we'll get the IP address from the
HOST, then "shell out" and run Vertica `admintools -t view_cluster -x` scrape
the results of the given node through its IP address to see if it's UP or DOWN.

## Running

To run the application use the following command:

    ./verticacheckd

There are various arguments that can be used to modify the address, port,
service-name, and HTTP read and write timeouts.

Current flag defaults:

    ./verticacheckd -h
      -address string
            HTTP address (default "127.0.0.1")
      -port int
            HTTP server listening port (default 3000)
      -service-name string
            Service name (default "verticacheckd")
      -timeouts duration
            HTTP Read and Write timeout (default 5s)

## Building

To build the project we have set a make task that'll only build Darwin and Linux
binaries for amd64.

    make release

## Tests

Running tests can be performed from the default make command or from the test
target

    make test

