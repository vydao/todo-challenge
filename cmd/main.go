package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	cors "github.com/rs/cors/wrapper/gin"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	migration, err := migrate.New("file://db/migration", "postgres://golang:secret@localhost:5432/todo-challenge?sslmode=disable")
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration successful")

	engine := gin.Default()
	engine.Use(cors.Default())
	groupV1 := engine.Group("/api/v1")
	groupV1.Handle(http.MethodGet, "/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": nil})
	})
	groupV1.Handle(http.MethodPost, "/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	log.Println(engine.Run(":8080"))
}
