package tests

import (
	"context"
	"testing"

	"github.com/IsaacDSC/pdv-golang/external/sqlc"
	"github.com/IsaacDSC/pdv-golang/src/infra/settings"
	"github.com/stretchr/testify/assert"
)

func TestSQLC(t *testing.T) {
	conn := settings.DbConn()
	tx, err := conn.Begin()
	assert.NoError(t, err)
	db := sqlc.New(tx)
	total, err := db.CounterClients(context.Background())
	tx.Commit()
	assert.NoError(t, err)
	assert.Equal(t, total, int64(0))
}
