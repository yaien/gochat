package sockets

import "github.com/yaien/gochat/auth"

type OutputMessage struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarURL"`
	Timestamp string `json:"timestamp"`
	Text      string `json:"text"`
}

type InputMessage struct {
	Message []byte
	User    *auth.User
}
