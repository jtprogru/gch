package uuids

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUUID(t *testing.T) {
	uuidStr := GetUUID()
	parsedUUID, err := uuid.Parse(uuidStr)

	// Check that no error occurred during parsing
	assert.NoError(t, err)

	// Check that the parsed UUID is valid
	assert.Equal(t, uuidStr, parsedUUID.String())
}
