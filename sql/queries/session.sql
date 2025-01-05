-- name: GetTokenById :one
select *
from bakery_session
where user_id = ?
;

-- name: GetTokenByToken :one
select *
from bakery_session
where token = ?
;
-- name: CreateToken :one
insert into bakery_session(
    user_id,
    token
)
values (?, ?)
RETURNING *;

-- name: UpdateToken :one
update bakery_session set
    token = ?
where user_id = ?
RETURNING *;

-- name: DeleteTokenById :exec
delete from bakery_session
where user_id = ?
;

-- name: DeleteTokenByToken :exec
delete from bakery_session
where token = ?
;

