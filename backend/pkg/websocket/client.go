package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client struct representing a websocket client
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Message struct representing a websocket message
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// Read allows client to read message from websocket
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message received: %+v\n", message)
	}
}
