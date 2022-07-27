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
	assert.Equal(t, result, true, "Input is 3 of kind cards in p1 hand, should be true")
}

func TestTongOneGameLose(t *testing.T) {
	result := IsPlayer1WinByTong([]string{"7D AS 5D AD 3C KC KS KA 5H 4S"})
	assert.Equal(t, result, true, "Input is 3 of kind cards in p2 hand, should be false")
}
