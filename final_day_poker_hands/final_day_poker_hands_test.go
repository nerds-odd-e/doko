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

func TestReturn1IfP1HasHighestCardIfHandIsSorted(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCard()}))
}

func TestReturn1IfP1WinsBySecondCard(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{"2S 3C 4D 5H 9D 2S 3D 4C 4H 9H"}))
}
func TestReturn1IfP1WinsByThirdCard(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{"2S 3C 4D 5H 9D 2S 3D 3C 5H 9H"}))
}

func TestReturn1IfP1WinsBylastCard(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{"3S 3C 4D 5H 6D 2S 3D 4C 5H 6H"}))
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
	assert.Equal(t, 1, pokerhands([]string{"2S 3C 4D 5H AH 2S 3D 4C 4H AS"}))
}

func makeP1HighestCardHighCardWithKing() string {
	return "2S 3C 4D 5H KH 2S 3D 4C 5H TD"
}

func TestReturn1IfP1HasHighestCardWithKing(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCardWithKing()}))
}

func TestSortTwoCards1(t *testing.T) {
	hand := Cards{"9H", "3C"}
	expected := Cards{"3C", "9H"}
	assert.Equal(t, expected, hand.sort())
}

func xTestSortTwoCards2(t *testing.T) {
	hand := Cards{"9H", "4C"}
	expected := Cards{"4C", "9H"}
	assert.Equal(t, expected, hand.sort())
}
