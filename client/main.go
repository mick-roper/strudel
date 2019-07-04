package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var server = flag.String("server", "ws://localhost:8080", "the server address")

func main() {
	flag.Parse()
	log.Println("client: connecting to", *server)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial(*server, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()

			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		// break out
		case <-done:
			return
		// on tick
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
			// on interrupt
		case <-interrupt:
			log.Println("interrupt received")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				log.Println("write close:", err)
				return
			}

			// wait for the socket to be closed
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
