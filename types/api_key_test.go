package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIKey(t *testing.T) {
	key := NewAPIKey(1, "default")

	assert.Equal(t, 32, len(key.Key))
}
