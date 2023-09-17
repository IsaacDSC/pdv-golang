-- name: CreateClient :exec
INSERT INTO client (
        client_id,
        email,
        username,
        telephone,
        created_at
    )
VALUES($1, $2, $3, $4, NOW());

-- name: CounterClients :one
SELECT COUNT(*) FROM client;