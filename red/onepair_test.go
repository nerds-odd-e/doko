package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsOnePair(t *testing.T) {
	assert.Equal(t, true, IsOnePair([]string{"AC", "AS", "2D", "3H", "5S"}))
	assert.Equal(t, true, IsOnePair([]string{"KC", "AS", "AD", "3H", "5S"}))
	assert.Equal(t, true, IsOnePair([]string{"KC", "KS", "AD", "3H", "5S"}))

	assert.Equal(t, false, IsOnePair([]string{"AC", "KS", "2D", "3H", "5S"}))
	assert.Equal(t, false, IsOnePair([]string{"TC", "AS", "2D", "3H", "5S"}))

}
