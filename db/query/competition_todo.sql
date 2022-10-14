-- name: CreateCompetitionTodo :one
INSERT INTO competition_todos (competition_id, todo_id)
VALUES ($1, $2)
RETURNING id, competition_id, todo_id, is_completed, created_at, updated_at;

-- name: GetCompetitionTodos :many
SELECT * FROM competition_todos WHERE competition_id = $1;

-- name: UpdateCompetitionTodo :one
UPDATE competition_todos SET is_completed = $1 WHERE id = $2 RETURNING *;