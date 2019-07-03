.PHONY: all

all: clean install test build

clean: clean-server clean-client

install: install-server install-client

build: build-server build-client

test: test-server test-client

# server
build-server:
	go build -o server/bin/strudel-server server/src/main.go

clean-server:
	rm -rf server/bin

install-server:
	go get -v ./server/...

test-server:
	go test ./server/...

# client
build-client:
	go build -o client/bin/strudel-client client/src/main.go

clean-client:
	rm -rf client/bin

install-client:
	go get -v ./client/...

test-server:
	go test ./client/...