package Tong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTongEmpty(t *testing.T) {
	result := IsPlayer1WinByTong([]string{})
	assert.Equal(t, result, false, "Input is empty game, should be false")
}

func TestTongOneGameWin(t *testing.T) {
	result := IsPlayer1WinByTong([]string{"KC KS KA 5H 4S 7D AS 5D AD 3C"})
	assert.Equal(t, result, false, "Input is 3 of kind cards in p1 hand, should be true")
}
