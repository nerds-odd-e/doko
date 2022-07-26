package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer1Win(t *testing.T) {
	type testcase struct {
		expected int
		input    []string
	}

	wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
	losecase := "3H 7H 6S 2C JS QH TD JC 2D 8S"
	cases := []testcase{
		{expected: 0, input: []string{}},
		{expected: 1, input: []string{wincase}},
		{expected: 0, input: []string{losecase}},
		{expected: 2, input: []string{wincase, wincase}},
		{expected: 2, input: []string{wincase, wincase, losecase}},
		{expected: 2, input: []string{losecase, wincase}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v game(s) Player 1 win %v time(s)", len(c.input), c.expected), func(t *testing.T) {
			assert.Equal(t, c.expected, Player1Win(c.input))
		})
	}
}
