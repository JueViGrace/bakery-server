-- name: GetSessionById :one
select *
from bakery_session
where id = ?
;

-- name: GetSessionByUser :many
select *
from bakery_session
where user_id = ?
;

-- name: GetSessionByUsername :many
select *
from bakery_session
where username = ?
;

-- name: CreateSession :exec
insert into bakery_session(
    id,
    refresh_token,
    access_token,
    username,
    user_id
)
values (?, ?, ?, ?, ?);

-- name: UpdateSession :exec
update bakery_session set
    refresh_token = ?,
    access_token = ?,
    username = ?
where id = ?;

-- name: DeleteSessionById :exec
delete from bakery_session
where id = ?
;

-- name: DeleteSessionByUser :exec
delete from bakery_session
where user_id = ?
;

-- name: DeleteSessionByToken :exec
delete from bakery_session
where refresh_token = ? or access_token = ?
;

