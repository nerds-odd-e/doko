package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Input_empty_File(t *testing.T) {
	result := Poker("empty.txt")
	assert.Equal(t, 0.0, result)
}
