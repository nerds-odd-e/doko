package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type PokerHandBuilder struct {
	highestHighcard string
}

func pokerHand() PokerHandBuilder {
	return PokerHandBuilder{}
}

func (b PokerHandBuilder) highCardOf(card string) PokerHandBuilder {
	b.highestHighcard = card
	return b
}

func (b PokerHandBuilder) please() string {
	return b.highestHighcard + "H 2C 3D 4S 6H"
}

func aGameOfTwoHighCards(p1HighCard string, p2HighCard string) []string {
	return []string{pokerHand().highCardOf(p1HighCard).please() + " " + pokerHand().highCardOf(p2HighCard).please()}
}

func twoGameOfTwoHighCards(p1HighCard []string, p2HighCard []string) []string {
	return []string{
		pokerHand().highCardOf(p1HighCard[0]).please() + " " + pokerHand().highCardOf(p2HighCard[0]).please(),
		pokerHand().highCardOf(p1HighCard[1]).please() + " " + pokerHand().highCardOf(p2HighCard[1]).please(),
	}
}

func TestPlayerWin0Game(t *testing.T) {
	a := []string{}
	assert.Equal(t, findWinnerPoker(a), 0)
}

func TestPlayerWin1Game(t *testing.T) {
	a := []string{
		"4H JH TH AH 3H 8H 9C 9D 7S 8H",
	}
	assert.Equal(t, findWinnerPoker(a), 1)
}
func TestPlayerWin2Game(t *testing.T) {
	a := []string{
		"4H JH TH AH 3H 8H 9C 9D 7S 8H",
		"4H JH TH AH 3H 8H 9C 9D 7S 8H",
	}
	assert.Equal(t, findWinnerPoker(a), 2)
}

func TestPlayerWin1in2Game(t *testing.T) {
	a := []string{
		"4H JH TH AH 3H 8H 9C 9D 7S 8H",
		"8H 9C 9D AS 9H 9H 9C 9D AS 9H",
	}
	assert.Equal(t, findWinnerPoker(a), 1)
}

func TestPlayerWin2in2Game(t *testing.T) {
	a := []string{
		"4H JH TH AH 3H 8H 9C 9D 7S 8H",
		"4H JH TH AH 3H 8H 9C 9D 7S 8H",
	}
	assert.Equal(t, findWinnerPoker(a), 2)
}

func TestP1WinByHighcard(t *testing.T) {
	assert.Equal(t, 1, findWinnerPoker(aGameOfTwoHighCards("A", "K")))
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("K", "Q")), 1)
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("Q", "J")), 1)
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("J", "T")), 1)
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("T", "9")), 1)
	assert.Equal(t, findWinnerPoker(twoGameOfTwoHighCards([]string{"T", "J"}, []string{"9", "T"})), 2)
}

func TestPlayer1LostByHighCard(t *testing.T) {
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("K", "A")), 0)
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("Q", "A")), 0)
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("J", "A")), 0)
	assert.Equal(t, findWinnerPoker(aGameOfTwoHighCards("T", "A")), 0)
	assert.Equal(t, findWinnerPoker(twoGameOfTwoHighCards([]string{"T", "J"}, []string{"A", "A"})), 0)
}
func TestPlayer1LostInTwoGameByHighCard(t *testing.T) {
	assert.Equal(t, findWinnerPoker(twoGameOfTwoHighCards([]string{"T", "T"}, []string{"A", "8"})), 1)
}

func xTestFindHighest(t *testing.T) {
	a := "4H JH TH AH 3H"
	assert.Equal(t, findHighest(a), "A")
	a = "4H JH AH 8H 3H"
	assert.Equal(t, findHighest(a), "A")
	a = "4H AH JH 8H 3H"
	assert.Equal(t, findHighest(a), "A")
	a = "AH 4H JH 8H 3H"
	assert.Equal(t, findHighest(a), "A")
	a = "QH 4H JH 8H 3H"
	assert.Equal(t, findHighest(a), "Q")
	a = "QH 4H JH KH 3H"
	assert.Equal(t, findHighest(a), "K")
	a = "TH 4H JH 9H 3H"
	assert.Equal(t, findHighest(a), "J")
	a = "2H 3H 4D 5S TD"
	assert.Equal(t, findHighest(a), "T")
	a = "2H 3H 4D 5S 9S"
	assert.Equal(t, findHighest(a), "9")
	a = "2H 3H 4D 5S 8S"
	assert.Equal(t, findHighest(a), "8")
	a = "2H 3H 4D 5S 7S"
	assert.Equal(t, findHighest(a), "7")

}

func TestFindWinner(t *testing.T) {

	content, err := ioutil.ReadFile("../../example_data/poker.txt")
	data := strings.Split(string(content), "\n")

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 380, findWinnerPoker(data))
}

func TestConvertPoint(t *testing.T) {
	assert.Equal(t, 10, ConvertPoint("A"))
	assert.Equal(t, 9, ConvertPoint("K"))
	assert.Equal(t, 8, ConvertPoint("Q"))
	assert.Equal(t, 7, ConvertPoint("J"))
	assert.Equal(t, 6, ConvertPoint("T"))
	assert.Equal(t, 5, ConvertPoint("9"))
	assert.Equal(t, 4, ConvertPoint("8"))
	assert.Equal(t, 3, ConvertPoint("7"))
}
