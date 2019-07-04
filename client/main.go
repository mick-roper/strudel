package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("", nil)

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

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalMessage, ""))

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
