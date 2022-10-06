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

	server := api.NewServer(db.NewStore(conn))
	engine := gin.Default()
	engine.Use(cors.Default())
	groupV1 := engine.Group("/api/v1")
	groupV1.Handle(http.MethodGet, "/users/:id", server.GetUserHandler)
	groupV1.Handle(http.MethodPost, "/users", server.CreateUserHandler)

	log.Println(engine.Run(":8080"))
}
