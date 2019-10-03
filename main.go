package main

import (
	"log"
	"net/http"
)

func main() {
	room := NewRoom()
	http.Handle("/", render("chat.html"))
	http.Handle("/room", room)
	go room.run()
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
