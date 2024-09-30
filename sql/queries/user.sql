-- name: GetUsers :many
SELECT *
FROM bakery_user;

-- name: GetUserById :one
SELECT *
FROM bakery_user
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM bakery_user
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO bakery_user (
        id,
        first_name,
        last_name,
        email,
        password,
        birth_date,
        phone,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
RETURNING *;

-- name: UpdateUser :one
UPDATE bakery_user
SET first_name = $1,
    last_name = $2,
    birth_date = $3,
    phone = $4,
    updated_at = NOW()
WHERE id = $5
RETURNING *;

-- name: DeleteUser :exec
UPDATE bakery_user
SET deleted_at = NOW()
WHERE id = $1;
