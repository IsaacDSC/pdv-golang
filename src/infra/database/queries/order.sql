-- name: CreateOrder :exec
INSERT INTO
    "order" (
        order_id,
        billing_type,
        price,
        discount,
        price_delivery,
        is_delivery,
        "status",
        client_id,
        delivery_id,
        created_at
    )
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, now());

-- name: UpdateStatusOrder :exec
UPDATE "order" SET status = $1, updated_at = now() WHERE order_id = $2;