package sockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/yaien/gochat/trace"
)

type Room struct {
	forward chan []byte
	join    chan *Client
	leave   chan *Client
	clients map[*Client]bool
	tracer  trace.Tracer
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("new client joined")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("client left")
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", string(msg))
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &Client{
		socket: socket,
		send:   make(chan []byte, 256),
		room:   r,
	}
	r.join <- client

	defer func() {
		r.leave <- client
	}()
	go client.write()
	client.read()
}

func NewRoom(tracer trace.Tracer) *Room {
	return &Room{
		forward: make(chan []byte),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
		tracer:  tracer,
	}
}
