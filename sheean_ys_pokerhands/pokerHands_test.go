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
	return "9H 3D 4D 5C 2C"
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

func HandsWithPlayer1Higher() string {
	return HighCardHandSorted() + " " + LowCardHandSorted()
}

func HandsWithPlayer2Higher() string {
	return LowCardHandSorted() + " " + HighCardHandSorted()
}
func TestTwoGames(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{HandsWithPlayer2Higher(), HandsWithPlayer2Higher()}))
	assert.Equal(t, 1, pokerHands([]string{HandsWithPlayer1Higher(), HandsWithPlayer2Higher()}))
	assert.Equal(t, 2, pokerHands([]string{HandsWithPlayer1Higher(), HandsWithPlayer1Higher()}))
}

func HighestCardSameRankThirdHigher() string {
	return "2C 3H 6D 8D 9S"
}

func HighestCardSameRankThirdLower() string {
	return "2C 3H 5D 8D 9S"
}

func TestTopTwoHighCardSameRank(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{HighestCardSameRankThirdHigher() + " " + HighestCardSameRankThirdLower()}))
}
