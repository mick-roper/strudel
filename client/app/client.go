package main

import (
	"flag"
	"log"

	"github.com/mick-roper/strudel/client/websocket"
)

var server = flag.String("server", "", "the strudel server")

func messageRecieved(message string) {
	log.Println(message)
}

func main() {
	flag.Parse()

	client, err := websocket.NewClient(*server, messageRecieved)

	if err != nil {
		log.Print(err)
	}

	defer client.Close()
}
