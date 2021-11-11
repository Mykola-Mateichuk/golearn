package main

import (
	"github.com/Mykola-Mateichuk/golearn/pkg/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Provide server functionality.
func main() {
	const pathPrefix = "/v1"

	r := mux.NewRouter()
	r.HandleFunc("/", server.IndexHandler)
	r.HandleFunc(pathPrefix + "/user", server.HandleCreateUser).Methods("POST")
	r.HandleFunc(pathPrefix + "/user/login", server.HandleLoginUser).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}