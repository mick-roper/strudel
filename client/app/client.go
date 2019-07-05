package main

import (
	"log"

	"github.com/mick-roper/strudel/client/websocket"
)

func main() {
	client, err := websocket.NewClient("abc")

	if err != nil {
		log.Print(err)
	}

	log.Print("%v", client)
}
