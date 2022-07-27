package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name     string
	expected float64
	input    []string
}

func TestPoker_winrate(t *testing.T) {
	cases := []TestCase{
		{name: "empty game", expected: 0.0, input: []string{""}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, WinRate(c.input))
		})
	}

}
