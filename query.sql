-- name: RegisterGoogle :exec
INSERT INTO users(
    name,email,oauth_provider,oauth_id,verified
    ) 
    VALUES(
        $1,$2,$3,$4,$5
    );

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUsersByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: RegisterManual :exec
INSERT INTO users(
    name,email,password,verified
)VALUES(
    $1,$2,$3,$4
);

