-- name: CreateDelivery :exec
INSERT INTO
    delivery (
        delivery_id,
        country,
        state,
        address,
        number_home,
        observation,
        token,
        client_id,
        created_at
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, NOW());

-- name: UpdateStatusDelivery :exec
UPDATE delivery SET status = $1, updated_at = NOW() WHERE delivery_id = $2;