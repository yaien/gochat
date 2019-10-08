package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yaien/gochat/auth"
	"github.com/yaien/gochat/sockets"
	"github.com/yaien/gochat/trace"
	"github.com/yaien/gochat/utils"
)

func chat(r *http.Request) interface{} {
	data := struct {
		User interface{}
	}{
		User: r.Context().Value("user"),
	}
	return data
}

func main() {
	auth.Setup()
	router := mux.NewRouter()
	tracer := trace.New(os.Stdout)
	room := sockets.NewRoom(tracer)
	router.Handle("/", auth.Auth(utils.Render("chat.html", chat)))
	router.Handle("/login", utils.Render("login.html", nil))
	router.HandleFunc("/auth/{provider}/login", auth.Login)
	router.HandleFunc("/auth/{provider}/callback", auth.Callback)
	router.Handle("/room", auth.Auth(room))
	go room.Run()
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
