package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	// 0.0 - 1.0 output
	// []string games
	//
	games := []string{}
	result := P1Winrate(games)

	assert.Equal(t, 1.0, result)
}

func TestHighCard(t *testing.T) {

}
func TestOpenEmptyFile_ReturnEmptyGame(t *testing.T) {
	assert.Equal(t, []string{}, OpenFile("data/empty_game.txt"))
}

func TestOpenOneLineFile_Return1Game(t *testing.T) {
	input := "data/one_line.txt"
	expected := []string{}

	assert.Equal(t, expected, OpenFile(input))
}
