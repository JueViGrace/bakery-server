-- name: GetSessionById :one
select *
from bakery_session
where user_id = ?
;

-- name: GetSessionByUsername :one
select *
from bakery_session
where username = ?
;

-- name: CreateSession :exec
insert into bakery_session(
    refresh_token,
    access_token,
    username,
    user_id
)
values (?, ?, ?, ?);

-- name: UpdateSession :exec
update bakery_session set
    refresh_token = ?,
    access_token = ?,
    username = ?
where user_id = ?;

-- name: DeleteSessionById :exec
delete from bakery_session
where user_id = ?
;

-- name: DeleteSessionByToken :exec
delete from bakery_session
where refresh_token = ? or access_token = ?
;

