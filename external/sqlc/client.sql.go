// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: client.sql

package sqlc

import (
	"context"
	"database/sql"
)

const counterClients = `-- name: CounterClients :one
SELECT COUNT(*) FROM client
`

func (q *Queries) CounterClients(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, counterClients)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createClient = `-- name: CreateClient :exec
INSERT INTO client (
        client_id,
        email,
        username,
        telephone,
        created_at
    )
VALUES($1, $2, $3, $4, NOW())
`

type CreateClientParams struct {
	ClientID  string
	Email     sql.NullString
	Username  sql.NullString
	Telephone sql.NullString
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) error {
	_, err := q.db.ExecContext(ctx, createClient,
		arg.ClientID,
		arg.Email,
		arg.Username,
		arg.Telephone,
	)
	return err
}
