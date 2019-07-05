package websocket

import (
	"errors"

	"github.com/gorilla/websocket"
)

// Client used to manage websocket connections
type Client struct {
	conn *websocket.Conn
}

// NewClient creates a new client
func NewClient(url string) (*Client, error) {
	if url == "" {
		return nil, errors.New("url must be provided")
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

// Send a message using the client
func (c *Client) Send(payload string) error {
	return errors.New("not implemented")
}
