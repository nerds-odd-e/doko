package poker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHighCard_Player1Win(t *testing.T) {
	answer :=  HighCardWin([]string{},[]string{})
	expected := false
	assert.Equal(t, expected, answer, "ต้องการ true ได้ %v", answer)
}