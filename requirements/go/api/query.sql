-- name: CreatePoll :exec
INSERT INTO Poll(id, title, createAt, updateAt)
VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- name: GetPoll :one
SELECT title FROM poll
WHERE id = $1;