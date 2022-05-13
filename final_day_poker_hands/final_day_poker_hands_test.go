package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//https://bicyclecards.com/how-to-play/basics-of-poker/
// c d h s
func TestReturn0WinsIfInputIsEmpty(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{}))
}

func makeP1HighestCardHighCard() string {
	return "2S 3C 4D 5H 9H 2S 3D 4C 5H 6D"
}

func makeUnsortedHand() string {
	return "9H 3C 4D 5H 2S"
}

func TestReturn1IfP1HasHighestCardIfHandIsSorted(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCard()}))
}

func TestReturn1IfP1WinsBySecondCard(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{"2S 3C 4D 5H 9D 2S 3D 4C 4H 9H"}))
}

func makeP2HighestCardHighCard() string {
	return "2S 3D 4C 5H 6D 2S 3C 4D 5H 9H"
}

func TestReturn0IfP2HasHighestCardIfHandIsSorted(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{makeP2HighestCardHighCard()}))
}

func xTestReturn1IfP1HasHighestCardAbove10IfHandIsSorted(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCardWithJack()}))
}

func makeP1HighestCardHighCardWithJack() string {
	return "2S 3C 4D 5H JH 2S 3D 4C 5H 6D"
}

func makeP1HighestCardHighCardWithAce() string {
	return "2S 3C 4D 5H AH 2S 3D 4C 5H JD"
}

func TestReturn1IfP1HasHighestCardWithAce(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCardWithAce()}))
}

func makeP1HighestCardHighCardWithKing() string {
	return "2S 3C 4D 5H KH 2S 3D 4C 5H 9D"
}

func TestReturn1IfP1HasHighestCardWithKing(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCardWithKing()}))
}
