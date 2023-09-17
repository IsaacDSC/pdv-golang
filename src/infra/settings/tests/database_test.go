package tests

import (
	"testing"

	"github.com/IsaacDSC/pdv-golang/src/infra/settings"
	"github.com/stretchr/testify/assert"
)

func TestConnectionDatabase(t *testing.T) {
	t.Run("Should be connect in database and create tables", func(t *testing.T) {
		db := settings.DbConn()
		err := db.Ping()
		assert.NoError(t, err)
	})
}
