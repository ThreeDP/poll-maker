-- name: CreatePoll :exec
INSERT INTO Poll(id, title, createAt, updateAt)
VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- name: GetPoll :one
SELECT title FROM poll
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO Users(Id, Name, Surname, Nickname, Email, Password, CreateAt)
VALUES($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP);