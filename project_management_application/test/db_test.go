package test

import (
	"project-management-application/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// This test will ensure that the ConnectDB function can connect to the database without errors.
	db.ConnectDB()

	// Check if DB is not nil
	assert.NotNil(t, db.DB)
}

// Add more tests for your database operations if needed.
