-- name: GetProducts :many
SELECT *
FROM bakery_product
WHERE bakery_product.deleted_at IS NULL;

-- name: GetProductById :one
SELECT *
FROM bakery_product
WHERE bakery_product.id = $1 AND
bakery_product.deleted_at IS NULL;

-- name: CreateProduct :one
INSERT INTO bakery_product (
        id,
        price,
        name,
        description,
        category,
        stock,
        image,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
RETURNING *;

-- name: UpdateProduct :one
UPDATE bakery_product
SET price = $1,
    name = $2,
    description = $3,
    category = $4,
    stock = $5,
    image = $6,
    updated_at = NOW()
WHERE id = $7
RETURNING *;

-- name: DeleteProduct :exec
UPDATE bakery_product
SET deleted_at = NOW()
WHERE id = $1;
