package sheenan_ys_pokerhands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func HighCardHandSorted() string {
	return "2C 3D 4D 5C 7C"
}
func LowCardHandSorted() string {
	return "2C 3H 4D 5D 6S"
}
func TestNoGames(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{}))
}

func HighCardHandUnSorted() string {
	return "2C 6H 4D 5D 9S"
}
func LowCardHandUnSorted() string {
	return "7C 3D 4D 5C 2C"
}
func TestPlayer1WinsWithHighCard(t *testing.T) {
	assert.Equal(t, 1, pokerHands([]string{HighCardHandSorted() + " " + LowCardHandSorted()}))
}
func TestPlayer2WinsWithHighCard(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{LowCardHandSorted() + " " + HighCardHandSorted()}))
}
func TestOneGameUnsorted(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{LowCardHandUnSorted() + " " + HighCardHandUnSorted()}))
	assert.Equal(t, 1, pokerHands([]string{HighCardHandUnSorted() + " " + LowCardHandUnSorted()}))
}

func HighestCardSameRankSecondHigher() string {
	return "2C 6H 4D 5D 9S"
}

func HighestCardSameRankSecondLower() string {
	return "9C 3D 4D 5C 2C"
}

func TestOneGameSameRank(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{HighestCardSameRankSecondLower() + " " + HighestCardSameRankSecondHigher()}))
	assert.Equal(t, 1, pokerHands([]string{HighestCardSameRankSecondHigher() + " " + HighestCardSameRankSecondLower()}))
}

func LowHandWithCourtCards() string {
	return "9C JD 4D KC TC"
}

func HighHandWithCourtCards() string {
	return "QC 6H 4D 5D AS"
}
func TestOneGameWithCourtCards(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{LowHandWithCourtCards() + " " + HighHandWithCourtCards()}))
	assert.Equal(t, 1, pokerHands([]string{HighHandWithCourtCards() + " " + LowHandWithCourtCards()}))
}

func TestOneGamme(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"2C 3D 4D 5C 6C 2C 3H 4D 5D 7S"}))
	assert.Equal(t, 1, pokerHands([]string{"2C 3D 4D 5C 7C 2C 3H 4D 5D 6S"}))
}

func TestTwoGames(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"9C JD 4D 5C 2C KC 6H 4D 5D 8S", "9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 1, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C", "9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 2, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C", "3S AD 4D 5C 2C 2C 6H QD 5D 9C"}))
}
