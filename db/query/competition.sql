-- name: CreateCompetition :one
INSERT INTO competitions (
  challenger_id,
  rival_id,
  challenge_id,
  status
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetCompetition :one
SELECT * FROM competitions WHERE id = $1;

-- name: GetCompetitionsByChallenge :many
SELECT * FROM competitions WHERE challenge_id = $1;

-- name: GetCompetitionsByUser :many
SELECT * FROM competitions WHERE challenger_id = $1;

-- name: GetCompetitionsByUserAndChallenge :one
SELECT * FROM competitions WHERE challenger_id = $1 AND rival_id = $2 AND challenge_id = $3 LIMIT 1;
