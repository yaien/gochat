package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yaien/gochat/sockets"
	"github.com/yaien/gochat/trace"
	"github.com/yaien/gochat/utils"
)

func main() {
	tracer := trace.New(os.Stdout)
	room := sockets.NewRoom(tracer)
	http.Handle("/", utils.Render("chat.html"))
	http.Handle("/room", room)
	go room.Run()
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
