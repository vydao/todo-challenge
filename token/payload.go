package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
	db "github.com/vydao/todo-challenge/db/sqlc"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(user db.User, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	payload := &Payload{
		ID:        tokenID,
		UserID:    user.ID,
		Username:  user.Username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, err
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
