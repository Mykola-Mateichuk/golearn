package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Id string
	UserName string
	password string
}

var users []User

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// handleCreateUser create user from request data.
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)
	uuidObj := uuid.New()
	user.Id = uuidObj.String()
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// handleLoginUser create login link.
func handleLoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	var id string
	for i := range users {
		if users[i].UserName == user.UserName && users[i].password == user.password {
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

// Provide server functionality.
func main() {
	const pathPrefix = "/v1"

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc(pathPrefix + "/user", handleCreateUser).Methods("POST")
	r.HandleFunc(pathPrefix + "/user/login", handleLoginUser).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}