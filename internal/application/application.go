package application

import (
	"encoding/json"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/models"
	"github.com/google/uuid"
	"net/http"
)

type UserServer struct {}

var users []models.User

func (server UserServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	uuidObj := uuid.New()
	user.Id = uuidObj.String()
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func (server UserServer) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

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
		fmt.Fprint(w, "Wrong user name or password")
	}
}