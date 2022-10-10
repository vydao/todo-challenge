package token

import (
	"time"

	db "github.com/vydao/todo-challenge/db/sqlc"
)

type Maker interface {
	CreateToken(user db.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
