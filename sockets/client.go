package sockets

import (
	"github.com/gorilla/websocket"
	"github.com/yaien/gochat/auth"
)

type Client struct {
	send   chan []byte
	socket *websocket.Conn
	room   *Room
	user   *auth.User
}

func (c *Client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- &InputMessage{
			Message: msg,
			User:    c.user,
		}
	}
}

func (c *Client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}
}
