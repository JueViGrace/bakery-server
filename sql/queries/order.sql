-- name: GetOrders :many
select *
from bakery_order
left join bakery_order_products on bakery_order.id = bakery_order_products.id
left join bakery_product on bakery_order_products.product_id = bakery_product.id
where bakery_order.deleted_at is null
;

-- name: GetOrdersByUser :many
select *
from bakery_order
left join bakery_order_products on bakery_order.id = bakery_order_products.id
left join bakery_product on bakery_order_products.product_id = bakery_product.id
where bakery_order.user_id = ? and bakery_order.deleted_at is null
;

-- name: GetOrderById :one
select *
from bakery_order
left join bakery_order_products on bakery_order.id = bakery_order_products.id
left join bakery_product on bakery_order_products.product_id = bakery_product.id
where bakery_order.id = ? and bakery_order.deleted_at is null
;

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
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: CreateOrderProducts :exec
INSERT INTO bakery_order_products(
    order_id,
    product_id,
    product_name,
    product_price,
    product_discount,
    product_rating,
    total_price,
    quantity
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateOrderStatus :exec
UPDATE bakery_order SET 
    status = ?,
    updated_at = ?
WHERE id = ?;

-- name: CancelOrder :exec
UPDATE bakery_order SET 
    status = "CANCELLED",
    updated_at = ?
WHERE id = ?;

