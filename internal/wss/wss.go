package wss

import (
	"github.com/Mykola-Mateichuk/golearn/internal/chat"
	"github.com/gorilla/websocket"
)


type WebsocketServer struct {
	Hub *chat.Hub // @todo change
	RequestUpgrader websocket.Upgrader
}

// NewWebsocketServer construct new object.
func NewWebsocketServer() WebsocketServer {
	var upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}

	hub := chat.NewHub()

	return WebsocketServer{Hub: hub, RequestUpgrader: upgrader}
}
