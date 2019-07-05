package websocket

import (
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

// Client used to manage websocket connections
type Client struct {
	conn *websocket.Conn
}

// NewClient creates a new client
func NewClient(url string, messageReceived func(string)) (*Client, error) {
	if url == "" {
		return nil, errors.New("url must be provided")
	}

	if messageReceived == nil {
		return nil, errors.New("no messageReceived function")
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			_, bytes, err := conn.ReadMessage()

			if err != nil {
				log.Print(err)
			}

			str := string(bytes)

			messageReceived(str)
		}
	}()

	return &Client{conn: conn}, nil
}

// Send a message using the client
func (c *Client) Send(payload string) error {
	return errors.New("not implemented")
}

// Close the client
func (c *Client) Close() {
	c.conn.Close()
}
