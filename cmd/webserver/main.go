// Package main provide functionality for webserver.
//
// By Mykola Mateichuk
package main

import (
	"github.com/Mykola-Mateichuk/golearn/internal/chatswaggeropenapi"
	"github.com/Mykola-Mateichuk/golearn/internal/handler"
	"github.com/Mykola-Mateichuk/golearn/internal/repository"
	"github.com/Mykola-Mateichuk/golearn/internal/service"
	"github.com/go-chi/chi"
	"net/http"
)

// Provide server functionality.
func main() {
	repo := repository.NewMemoryStorage()
	userservice := service.UserService{
		Repo: repo,
	}
	u := handler.UserServer{
		Userservice: userservice,
	}
	h := chatswaggeropenapi.HandlerFromMuxWithBaseURL(u, chi.NewRouter(), "/v1")
	http.ListenAndServe(":8080", h)
}