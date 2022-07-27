package poker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHighCard_Player1Win(t *testing.T) {
	answer :=  HighCardWin([]string{"6D", "2D", "2S", "2H", "2C"},[]string{"5D", "2D", "2S", "2H", "2C"})
	expected := true
	assert.Equal(t, expected, answer, "ต้องการ true ได้ %v", answer)
}

func TestHighCard_Player1Lose(t *testing.T) {
	answer :=  HighCardWin([]string{"4D", "1D", "2S", "2H", "2C"},[]string{"7D", "2D", "2S", "2H", "2C"})
	expected := false
	assert.Equal(t, expected, answer, "ต้องการ false ได้ %v", answer)
}