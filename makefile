.PHONY: all

all: clean install test build

install: install-server install-client

clean: clean-server clean-client

build: build-server build-client

test: test-server test-client

# server
build-server:
	cd server/app && go build -o ../../bin/strudel-server server.go

clean-server:
	rm -rf server/bin

install-server:
	cd server && go mod vendor

test-server:
	echo "no tests"

# client
build-client:
	cd client/app && go build -o ../../bin/strudel-client client.go

clean-client:
	rm -rf client/bin

install-client:
	cd client/app && go mod vendor

test-client:
	echo "no tests"
