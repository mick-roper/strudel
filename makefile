.PHONY: all

all: clean install test build

clean: clean-server clean-client

install: install-server install-client

build: build-server build-client

test: test-server test-client

# server
build-server:
	go build -o server/bin/strudel-server server/main.go

clean-server:
	rm -rf server/bin

install-server:
	echo "nothing to install"

test-server:
	echo "no tests"

# client
build-client:
	go build -o client/bin/strudel-client client/main.go

clean-client:
	rm -rf client/bin

install-client:
	cd client go get

test-client:
	echo "no tests"
