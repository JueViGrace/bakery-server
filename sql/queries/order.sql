-- name: GetOrders :many
SELECT *
FROM bakery_order
    LEFT JOIN bakery_order_products ON bakery_order.id = bakery_order_products.id
    LEFT JOIN bakery_product ON bakery_order_products.product_id = bakery_product.id
WHERE bakery_order.deleted_at IS NULL;

-- name: GetOrdersByUser :many
SELECT *
FROM bakery_order
    LEFT JOIN bakery_order_products ON bakery_order.id = bakery_order_products.id
    LEFT JOIN bakery_product ON bakery_order_products.product_id = bakery_product.id
WHERE bakery_order.user_id = $1
    AND bakery_order.deleted_at IS NULL;

-- name: GetOrderById :one
SELECT *
FROM bakery_order
    LEFT JOIN bakery_order_products ON bakery_order.id = bakery_order_products.id
    LEFT JOIN bakery_product ON bakery_order_products.product_id = bakery_product.id
WHERE bakery_order.id = $1
    AND bakery_order.deleted_at IS NULL;

-- name: CreateOrder :exec
INSERT INTO bakery_order (
        id,
        total_amount,
        payment_method,
        status,
        user_id,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, NOW(), NOW());

-- name: CreateOrderProducts :exec
INSERT INTO bakery_order_products(
        order_id,
        product_id,
        price,
        quantity
    )
VALUES ($1, $2, $3, $4);

-- name: UpdateOrderStatus :exec
UPDATE bakery_order
SET status = $1,
    updated_at = NOW()
WHERE id = $2;

-- name: DeleteOrder :exec
UPDATE bakery_order
SET status = $1,
    deleted_at = NOW()
WHERE id = $2;
