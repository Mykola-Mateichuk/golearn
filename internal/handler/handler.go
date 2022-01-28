// Package handler handle requests.
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/chat"
	"github.com/Mykola-Mateichuk/golearn/internal/chatswaggeropenapi"
	"github.com/Mykola-Mateichuk/golearn/internal/model"
	"github.com/Mykola-Mateichuk/golearn/internal/service"
	"github.com/Mykola-Mateichuk/golearn/internal/wss"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// UserServer contain all needed dependencies.
type UserServer struct {
	Userservice service.UserService
	WebsocketServer wss.WebsocketServer
}

// CreateUser create a new user.
func (server UserServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user model.User

	decErr := json.NewDecoder(r.Body).Decode(&user)
	if decErr != nil {
		fmt.Fprint(w, "Can't decode user")
		return
	}

	uuidObj := uuid.New()
	user.Id = uuidObj.String()
	user, err := server.Userservice.AddUser(user)

	if err != nil {
		panic(err)
		//fmt.Fprint(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// LoginUser create link if user exist.
func (server UserServer) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	decErr := json.NewDecoder(r.Body).Decode(&user)
	if decErr != nil {
		fmt.Fprint(w, "Can't decode user")
		return
	}

	url, err := server.Userservice.GetLoginLink(user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if url != "" {
		fmt.Fprintf(w, "You link to login into chat: %s", url)
		return
	} else {
		fmt.Fprint(w, "Wrong user name or password")
	}
}

// WsRTMStart Start websocket connection.
func (server UserServer) WsRTMStart(w http.ResponseWriter, r *http.Request, params chatswaggeropenapi.WsRTMStartParams) {
	ctx := r.Context()
	userName := ctx.Value("UserName").(string)

	server.WebsocketServer.RequestUpgrader.CheckOrigin = func(r *http.Request) bool {return true}
	conn, err := server.WebsocketServer.RequestUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	user, err := server.Userservice.GetUserByName(userName)
	if err != nil {
		log.Print("Can't get user:", err)
		return
	}
	client := &chat.Client{Hub: server.WebsocketServer.Hub, Conn: conn, Send: make(chan []byte, 256), User: user}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}

// GetActiveUsers create list with active users.
func (server UserServer) GetActiveUsers(w http.ResponseWriter, r *http.Request) {
	ausers := ""
	aclients := server.WebsocketServer.Hub.GetClients()

	for client := range aclients {
		ausers = ausers + ", " + client.User.UserName
	}
	fmt.Fprintf(w, "Active users: %s", ausers)
}
