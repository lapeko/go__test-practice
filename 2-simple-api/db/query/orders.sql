-- name: CreateOrder :one
INSERT INTO orders (user_id, product_name, amount, price)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderById :one
SELECT * FROM orders WHERE id = $1;

-- name: GetOrdersByUserId :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: UpdateOrderStatusById :one
UPDATE orders
set status = $2, updated_at = now()
WHERE id = $1
RETURNING *;
