package main

import (
	"log"

	"github.com/mick-roper/strudel/client/websocket"
)

func messageRecieved(message string) {
	log.Println(message)
}

func main() {
	client, err := websocket.NewClient("abc", messageRecieved)

	if err != nil {
		log.Print(err)
	}

	defer client.Close()
}
