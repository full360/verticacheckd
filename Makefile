VERSION=0.1.0
GO_LDFLAGS="-X main.version=$(VERSION)"

default: test

install:
	@go get -u ./...

test:
	@go test -v -race -cover .

bin:
	@mkdir -p bin
	@rm -rf bin/*

release: release-darwin \
	release-linux

release-darwin: bin
	GOOS=darwin GOARCH=amd64 go build -ldflags=$(GO_LDFLAGS) -o bin/verticacheckd ./cmd/verticacheckd
	cd bin && tar -cvzf verticacheckd$(VERSION).darwin-amd64.tgz verticacheckd
	rm bin/verticacheckd

release-linux: bin
	GOOS=linux GOARCH=amd64 go build -ldflags=$(GO_LDFLAGS) -o bin/verticacheckd ./cmd/verticacheckd
	cd bin && tar -cvzf verticacheckd$(VERSION).linux-amd64.tgz verticacheckd
	rm bin/verticacheckd

.PHONY: default install test bin release
