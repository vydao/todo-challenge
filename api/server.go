package api

import db "github.com/vydao/todo-challenge/db/sqlc"

type Server struct {
	store db.Store
}

func NewServer(store db.Store) *Server {
	return &Server{
		store: store,
	}
}
