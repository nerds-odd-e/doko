package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoker_empty_game(t *testing.T) {
	t.Run("empty game", func(t *testing.T) {
		assert.Equal(t, 0.0, WinRate([]string{""}))
	})

}
