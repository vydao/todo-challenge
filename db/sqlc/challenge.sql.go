// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: challenge.sql

package db

import (
	"context"
	"time"
)

const createChallenge = `-- name: CreateChallenge :one
INSERT INTO challenges (
  start_date,
  name,
  description,
  user_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, start_date, name, description, user_id, created_at, updated_at
`

type CreateChallengeParams struct {
	StartDate   time.Time `json:"start_date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      int64     `json:"user_id"`
}

func (q *Queries) CreateChallenge(ctx context.Context, arg CreateChallengeParams) (Challenge, error) {
	row := q.db.QueryRowContext(ctx, createChallenge,
		arg.StartDate,
		arg.Name,
		arg.Description,
		arg.UserID,
	)
	var i Challenge
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.Name,
		&i.Description,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getChallenge = `-- name: GetChallenge :one
SELECT id, start_date, name, description, user_id, created_at, updated_at FROM challenges WHERE id = $1
`

func (q *Queries) GetChallenge(ctx context.Context, id int64) (Challenge, error) {
	row := q.db.QueryRowContext(ctx, getChallenge, id)
	var i Challenge
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.Name,
		&i.Description,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
