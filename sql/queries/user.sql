-- name: GetUsers :many
select *
from bakery_user
;

-- name: GetUserById :one
select *
from bakery_user
where id = ?
;

-- name: GetUserByEmail :one
select *
from bakery_user
where email = ?
;

-- name: CreateUser :one
INSERT INTO bakery_user (
    id,
    first_name,
    last_name,
    username,
    email,
    password,
    phone_number,
    birth_date,
    address1,
    address2,
    gender,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateUser :one
UPDATE bakery_user SET
    first_name = ?,
    last_name = ?,
    phone_number = ?,
    birth_date = ?,
    address1 = ?,
    address2 = ?,
    gender = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: UpdateEmail :one
UPDATE bakery_user SET
    email = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: UpdateUsername :one
UPDATE bakery_user SET
    username = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
UPDATE bakery_user SET 
    deleted_at = ?
WHERE id = ?;

