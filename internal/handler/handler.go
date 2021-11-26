// Package handler handle requests.
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/model"
	"github.com/Mykola-Mateichuk/golearn/internal/service"
	"github.com/google/uuid"
	"net/http"
)

// UserServer contain all needed dependencies.
type UserServer struct {
	Userservice service.UserService
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
		fmt.Fprint(w, err)
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