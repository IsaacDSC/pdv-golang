package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConnectionDatabase(t *testing.T) {
	conn := GetConnDatabaseStore()
	assert.NoError(t, conn.Error)
}
