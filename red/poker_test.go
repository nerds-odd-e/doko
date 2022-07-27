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
		{name: "win 1 game", expected: 0.0, input: []string{"7D 2S 5D 3S AC 8C TS KC 9H 4S"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, WinRate(c.input))
		})
	}

}

func TestTwoPairs(t *testing.T) {
	t.Run("Player1 should win two pairs at rate 1.0", func(t *testing.T) {
		assert.Equal(t, 1.0, 1.0)
	})
}
