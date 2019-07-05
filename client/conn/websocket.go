package conn

import "github.com/gorilla/websocket"

// Client that manages websocket connections
type Client struct {
	conn websocket.Conn
}

// NewClient creates a new client
func NewClient(url string) (*Client, error) {
	c, _, err := websocket.DefaultDialler.Dial(url, nil)

	if err != nil {
		return nil, err
	}

	return &Client{ 
		conn: c,
	}, nil
}
