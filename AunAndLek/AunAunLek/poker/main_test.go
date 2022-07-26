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

	cases := []testcase{
		{expected: 0, input: []string{}},
		{expected: 1, input: []string{"3H 7H 6S KC JS QH TD JC 2D 8S"}},
		{expected: 2, input: []string{"AH 7H 6S KC JS QH TD JC 2D KS", "TH 7H 6S KC JS QH TD JC 2D 9S"}},
		{expected: 2, input: []string{"AH 7H 6S KC JS QH TD JC 2D KS", "TH 7H 6S KC JS QH TD JC 2D 9S", "3H 7H 6S 2C JS QH TD JC 2D 8S"}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v game Player 1 win %v time(s)", len(c.input), c.expected), func(t *testing.T) {
			assert.Equal(t, c.expected, Player1Win(c.input))
		})
	}
}
