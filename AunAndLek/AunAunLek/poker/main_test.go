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
	}

	expected := cases[0].expected
	input := cases[0].input
	t.Run(fmt.Sprintf("%v game Player 1 win %v time(s)", len(input), expected), func(t *testing.T) {
		assert.Equal(t, expected, Player1Win(input))
	})
}
