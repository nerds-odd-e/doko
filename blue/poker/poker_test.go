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

	t.Run("P1 win", func(t *testing.T) {
		game := "AC 6S 5D 3C 2C KS 5C 4H 3D 2C"
		assert.Equal(t, true, P1Win(game))
	})

	t.Run("P2 wins", func(t *testing.T) {
		game := "KS 5C 4H 3D 2C AC 6S 5D 3C 2C"
		assert.Equal(t, true, P1Win(game))
	})
}
func TestOpenEmptyFile_ReturnEmptyGame(t *testing.T) {
	assert.Equal(t, []string{}, OpenFile("data/empty_game.txt"))
}

func TestOpenOneLineFile_Return1Game(t *testing.T) {
	input := "data/one_line.txt"
	expected := []string{"AA AA AA AA AA AA AA AA AA AA"}

	assert.Equal(t, expected, OpenFile(input))
}
