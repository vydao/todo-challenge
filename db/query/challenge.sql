-- name: CreateChallenge :one
INSERT INTO challenges (
  start_date,
  name,
  description,
  user_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetChallenge :one
SELECT * FROM challenges WHERE id = $1;
