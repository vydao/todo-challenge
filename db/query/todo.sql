-- name: CreateTodo :one
INSERT INTO todos (
  name,
  challenge_id,
  period,
  point
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: GetTodosByChallenge :many
SELECT * FROM todos WHERE challenge_id = $1;