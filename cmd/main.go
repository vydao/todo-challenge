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

var connStr = "postgres://hsiisitqhqcjla:799464783d372b35e0c02fa1379b98166268d83599f05ad75aa5304ede6800a0@ec2-52-70-45-163.compute-1.amazonaws.com:5432/d4sq4jri2g0fsl"

func main() {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	migration, err := migrate.New("file://db/migration", connStr)
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
	apiV1.Handle(http.MethodPost, "/users", server.CreateUserHandler)

	authV1 := apiV1.Group("/")
	authV1.Use(api.AuthMiddleWare(tokenMaker))
	authV1.Handle(http.MethodGet, "/users/:id", server.GetUserHandler)
	authV1.Handle(http.MethodPost, "/challenges", server.CreateChallengeHandler)
	authV1.Handle(http.MethodPost, "/challenges/:challenge_id/todos", server.CreateTodoHandler)
	authV1.Handle(http.MethodGet, "/challenges/:challenge_id/todos", server.GetTodosByChallengeHandler)
	authV1.Handle(http.MethodPost, "/challenges/:challenge_id/accept", server.AcceptChallengeHandler)

	log.Println(engine.Run(":8080"))
}
