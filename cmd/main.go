package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/vydao/todo-challenge/api"
	db "github.com/vydao/todo-challenge/db/sqlc"
	"github.com/vydao/todo-challenge/token"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	conn, err := sql.Open("postgres", "postgresql://golang:secret@localhost:5432/todo-challenge?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	migration, err := migrate.New("file://db/migration", "postgres://golang:secret@localhost:5432/todo-challenge?sslmode=disable")
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration successful")

	store := db.NewStore(conn)
	tokenMaker, err := token.NewJWTMaker("loremipsumdolorsitamet1234567891012131343")
	if err != nil {
		log.Fatal("Cannot init token maker:", err)
	}
	server := api.NewServer(store, tokenMaker)

	engine := gin.Default()
	engine.Use(cors.Default())
	apiV1 := engine.Group("/api/v1")
	apiV1.Handle(http.MethodPost, "/users/login", server.LoginUserHandler)

	authV1 := apiV1.Group("/")
	authV1.Use(api.AuthMiddleWare(tokenMaker))
	authV1.Handle(http.MethodGet, "/users/:id", server.GetUserHandler)
	authV1.Handle(http.MethodPost, "/users", server.CreateUserHandler)
	authV1.Handle(http.MethodPost, "/challenges", server.CreateChallengeHandler)
	authV1.Handle(http.MethodPost, "/challenges/:challenge_id/todos", server.CreateTodoHandler)
	authV1.Handle(http.MethodGet, "/challenges/:challenge_id/todos", server.GetTodosByChallengeHandler)

	log.Println(engine.Run(":8080"))
}
