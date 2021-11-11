// Package server provide functionality to dial with server functionality.
//
// By Mykola Mateichuk
package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id string
	UserName string	`validate:"required"`
	Password string	`validate:"required"`
}

var users []User

// IndexHandler responds to requests with our greeting.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// HandleCreateUser create user from request data.
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)

	// Validate user.
	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Wrong user name or password")
		return
	}

	// Add new user to list.
	uuidObj := uuid.New()
	user.Id = uuidObj.String()
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// HandleLoginUser create login link.
func HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Validate user.
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Wrong user name or password")
		return
	}

	// Find user with name and id.
	var id string
	for i := range users {
		if users[i].UserName == user.UserName && users[i].Password == user.Password {
			id = users[i].Id
		}
	}

	if id != "" {
		url := "ws://fancy-chat.io//ws&token=one-time-token&id=" + id
		fmt.Fprintf(w, "You link to login into chat: %s", url)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "There are no user with this user name and password.")
	}
}