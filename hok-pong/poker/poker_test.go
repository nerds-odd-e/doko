package poker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/hok-pong/poker"
)

func TestPoker_FirstHighCard_Player1_Win(t *testing.T) {
	data := "KC 5C 4S 3H 2S QD 5S 4D 3S 2C"
	actual := poker.Winner(data)
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestPoker_FirstHighCard_Player2_Win(t *testing.T) {
	data := "QD 5C 4S 3H 2S KC 5S 4D 3S 2C"
	actual := poker.Winner(data)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestPoker_FullHandHighCard_Player2_Win(t *testing.T) {
	data := "QD 5C 4S 3H 2S QH 6S 4D 3S 2C"
	actual := poker.Winner(data)
	expected := 2
	assert.Equal(t, expected, actual)
}
