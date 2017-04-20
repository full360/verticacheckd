# Vertica State

Vertica State is a web service that returns the state (up/down) of the Vertica
node it's running on.

## Detail

It will receive a request from an AWS ELB that will contain an IP address, then
"shell out" and run Vertica `admintools -t view_cluster -x` scrape the results
of the given node through its IP address to see if it's UP or DOWN.
