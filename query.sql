-- name: InsertUsers :exec
INSERT INTO users(name,email,password) VALUES($1,$2,$3);

-- name: GetUsers :many
SELECT * FROM users;