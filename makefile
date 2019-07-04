.PHONY: all

all: clean install test build

install: install-server install-client

clean: clean-server clean-client

build: build-server build-client

test: test-server test-client

# server
build-server:
	cd server && go build -o bin/strudel-server main.go

clean-server:
	rm -rf server/bin

install-server:
	echo "nothing to instal"

test-server:
	echo "no tests"

# client
build-client:
	cd client && go build -o bin/strudel-client main.go

clean-client:
	rm -rf client/bin

install-client:
	cd client && go mod vendor

test-client:
	echo "no tests"
