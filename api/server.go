package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vydao/todo-challenge/db/sqlc"
	"github.com/vydao/todo-challenge/token"
)

type Server struct {
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(store db.Store, tokenMaker token.Maker) *Server {
	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
