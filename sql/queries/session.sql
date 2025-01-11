-- name: GetSessionById :one
select *
from bakery_session
where user_id = ?
;

-- name: GetSessionByToken :one
select *
from bakery_session
where token = ?
;

-- name: CreateSession :one
insert or replace into bakery_session(
    user_id,
    token
)
values (?, ?)
RETURNING *;

-- name: DeleteSessionById :exec
delete from bakery_session
where user_id = ?
;

-- name: DeleteSessionByToken :exec
delete from bakery_session
where token = ?
;

