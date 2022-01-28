// Package main provide functionality for webserver.
//
// By Mykola Mateichuk
package main

import (
	"database/sql"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/chatswaggeropenapi"
	"github.com/Mykola-Mateichuk/golearn/internal/handler"
	"github.com/Mykola-Mateichuk/golearn/internal/middleware"
	"github.com/Mykola-Mateichuk/golearn/internal/repository"
	"github.com/Mykola-Mateichuk/golearn/internal/service"
	"github.com/Mykola-Mateichuk/golearn/internal/token"
	"github.com/Mykola-Mateichuk/golearn/internal/wss"
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

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewPostgreSqlStorage(db)
	tokenMaker, err := token.NewPasetoMaker("01234567890123456789012345678912")
	userservice := service.NewUserService(repo, tokenMaker)
	websocketServer := wss.NewWebsocketServer()
	go websocketServer.Hub.Run()

	u := handler.UserServer{
		Userservice: userservice,
		WebsocketServer: websocketServer,
	}
	//h := chatswaggeropenapi.HandlerFromMuxWithBaseURL(u, chi.NewRouter(), "/v1")

	var middlewareFuncs []chatswaggeropenapi.MiddlewareFunc
	middlewareFuncs = append(middlewareFuncs, middleware.MiddlewareLogAllErrors)
	middlewareFuncs = append(middlewareFuncs, middleware.MiddlewareLogAllCalls)
	middlewareFuncs = append(middlewareFuncs, middleware.MiddlewareLogPanicsAndRecover)
	middlewareFuncs = append(middlewareFuncs, middleware.AuthMiddleware)

	options := chatswaggeropenapi.ChiServerOptions{
		BaseURL: "/v1",
		BaseRouter: chi.NewRouter(),
		Middlewares: middlewareFuncs,
		ErrorHandlerFunc: middleware.ErrorHandler,
	}
	r := chatswaggeropenapi.HandlerWithOptions(u, options)


	http.ListenAndServe(":8080", r)
}