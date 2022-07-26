package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirstPlayerWinCount(t *testing.T) {
	result := getFirstPlayerWinCount("./files/empty.txt")
	assert.Equal(t, result, 0, "Input is empty file which is no game, should get 0")
}
