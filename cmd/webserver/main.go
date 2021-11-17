// Package main provide functionality for webserver.
//
// By Mykola Mateichuk
package main

import (
	"github.com/Mykola-Mateichuk/golearn/internal/application"
	"github.com/Mykola-Mateichuk/golearn/internal/server"
	"github.com/go-chi/chi"
	"net/http"
)

// Provide server functionality.
func main() {
	u := application.UserServer{}
	h := server.HandlerFromMuxWithBaseURL(u, chi.NewRouter(), "/v1")
	http.ListenAndServe(":8080", h)
}