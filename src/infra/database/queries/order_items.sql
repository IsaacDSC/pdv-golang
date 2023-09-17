-- name: CreateOrderItems :exec
INSERT INTO
    order_items (
        order_items_id,
        product_id,
        quantity,
        price,
        created_at,
        order_id
    )
VALUES
($1, $2, $3, $4, $5,now());