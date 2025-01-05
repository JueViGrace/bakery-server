-- name: GetProducts :many
select *
from bakery_product
where bakery_product.deleted_at is null
;

-- name: GetProductById :one
select *
from bakery_product
where bakery_product.id = ? and bakery_product.deleted_at is null
;

-- name: CreateProduct :one
INSERT INTO bakery_product (
    id,
    name,
    description,
    category,
    price,
    stock,
    issued,
    has_stock,
    discount,
    rating,
    images,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateProduct :one
UPDATE bakery_product SET
    description = ?,
    category = ?,
    price = ?,
    stock = ?,
    issued = ?,
    has_stock = ?,
    discount = ?,
    rating = ?,
    images = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteProduct :exec
UPDATE bakery_product SET 
    deleted_at = ?
WHERE id = ?;

