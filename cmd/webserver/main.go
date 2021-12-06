// Package main provide functionality for webserver.
//
// By Mykola Mateichuk
package main

import (
	"database/sql"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/chatswaggeropenapi"
	"github.com/Mykola-Mateichuk/golearn/internal/handler"
	"github.com/Mykola-Mateichuk/golearn/internal/repository"
	"github.com/Mykola-Mateichuk/golearn/internal/service"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// Provide server functionality.
func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres_docker"
		dbname   = "postgres"
	)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	//db, err := sql.Open("postgres", "postgresql://postgres:postgres@postgres/chat")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewPostgreSqlStorage(db)
	userservice := service.NewUserService(repo)
	u := handler.UserServer{
		Userservice: userservice,
	}
	h := chatswaggeropenapi.HandlerFromMuxWithBaseURL(u, chi.NewRouter(), "/v1")
	http.ListenAndServe(":8080", h)
}