package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsOnePair(t *testing.T) {
	assert.Equal(t, true, IsOnePair([]string{"AC", "AS", "2D", "3H", "5S"}))
}
