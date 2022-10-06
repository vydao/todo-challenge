package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	cors "github.com/rs/cors/wrapper/gin"
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

	store := db.NewStore(conn)

	engine := gin.Default()
	engine.Use(cors.Default())
	groupV1 := engine.Group("/api/v1")
	groupV1.Handle(http.MethodGet, "/users/:id", func(ctx *gin.Context) {
		var req GetUserRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		user, err := store.GetUser(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.AbortWithError(http.StatusNotFound, err)
				return
			}
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": GetUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Score:    user.Score,
		}})
	})
	groupV1.Handle(http.MethodPost, "/users", func(ctx *gin.Context) {
		userReq := db.CreateUserParams{}
		if err := ctx.BindJSON(&userReq); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		_, err := store.CreateUser(ctx, userReq)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	})

	log.Println(engine.Run(":8080"))
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetUserResponse struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}
