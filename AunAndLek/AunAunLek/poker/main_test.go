package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer1Win(t *testing.T) {
	t.Run("Player 1 win with highest card A", func(t *testing.T) {
		input := "AC 2C TC JC 5C KC JC 6C 2C QC"
		expected := true

		result := IsP1Win(input)

		assert.Equal(t, expected, result)
	})

	t.Run("Player 2 win with highest card A", func(t *testing.T) {
		input := "KC 2C TC JC 5C AC JC 6C 2C QC"
		expected := false

		result := IsP1Win(input)

		assert.Equal(t, expected, result)
	})

	// t.Run("Player 1 win with highest card K", func(t *testing.T) {
	// 	input := "KC 2C TC JC 5C TC JC 6C 2C QC"
	// 	expected := true

	// 	result := IsP1Win(input)

	// 	assert.Equal(t, expected, result)
	// })
}

// func TestPlayer1WinCount(t *testing.T) {
// 	type testcase struct {
// 		expected int
// 		input    []string
// 	}

// 	wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
// 	losecase := "3H 7H 6S 2C JS QH TD JC 2D 8S"
// 	cases := []testcase{
// 		{expected: 0, input: []string{}},
// 		{expected: 1, input: []string{wincase}},
// 		{expected: 0, input: []string{losecase}},
// 		{expected: 2, input: []string{wincase, wincase}},
// 		{expected: 2, input: []string{wincase, wincase, losecase}},
// 		{expected: 1, input: []string{losecase, wincase}},
// 		{expected: 2, input: []string{losecase, wincase, wincase}},
// 	}

// 	for _, c := range cases {
// 		t.Run(fmt.Sprintf("%v game(s) Player 1 win %v time(s)", len(c.input), c.expected), func(t *testing.T) {
// 			assert.Equal(t, c.expected, Player1WinCount(c.input))
// 		})
// 	}
// }
