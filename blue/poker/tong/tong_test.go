package Tong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTong(t *testing.T) {
	result := IsPlayer1WinByTong([]string{})
	assert.Equal(t, result, false, "Should be false")
}
