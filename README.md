# verticacheckd

verticacheckd is a web service that returns the state of the Vertica node it's
running on by returning an HTTP 200 or 500 depending if it's UP or in any
other state.

In detail, It will receive a request from an AWS ELB and we'll get the IP
address from the HOST, then "shell out" and run Vertica
`/opt/vertica/bin/admintools -t view_cluster -x` scrape the results of the given
node through its IP address to see if it's UP or DOWN, and return a response
through the endpoint that could either be HTTP 200 or 500.

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
      -timeout duration
            HTTP Read and Write timeout (default 5s)

Current state endpoints:

- `{service-name}/state` this will return a global state for the node
- `{service-name}/dbs/{name}/state` will return the state of a specific db given
  a name

## Installing for local development

First make sure you have the `GOPATH` set up. If that's the case clone the
project inside the `GOPATH`

    git clone git@gitlab.full360.com:full360/verticacheckd.git $GOPATH/src/gitlab.full360.com/full360/verticacheckd

Then install dependencies:

    make install

OR:

    go get -u ./...

with that you are done.

## Building and Releasing

To build the project we have set a make task that'll only build Darwin and Linux
binaries for amd64. Remember to Bump the version inside the `Makefile` before
releasing and that's about it.

    make release

## Tests

Running tests can be performed from the default make command or from the test
target

    make test

