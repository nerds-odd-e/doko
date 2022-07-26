package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer1Win(t *testing.T) {
	t.Run("0 game Player 1 win 0 time(s)", func(t *testing.T) {
		expected := 0
		input := [][]string{}

		assert.Equal(t, expected, Player1Win(input))
	})
}
