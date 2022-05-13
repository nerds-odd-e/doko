package final_day_poker_hands

import (
	"sort"
	"strconv"
	"strings"
)

func pokerhands(games []string) int {
	p1WinCount := 0
	for _, game := range games {
		cards := strings.Split(game, " ")
		if P1WinsCompareHighCard(cards) || P1WinsOnePair(cards) {
			p1WinCount += 1
		}

	}
	return p1WinCount
}

func getFaceValue(face string) int {
	switch face {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	}
	faceValue, _ := strconv.Atoi(face)
	return faceValue
}

func P1WinsOnePair(cards []string) bool {
	if cards[0][0] == cards[1][0] {
		return true
	}
	return cards[3][0] == cards[4][0]
}

type Hand []string

func (hand1 Hand) isBiggerThan(hand2 Hand) bool {
	for i := 4; i >= 0; i-- {
		if hand1[i][:1] != hand2[i][:1] {
			return getFaceValue(hand1[i][:1]) > getFaceValue(hand2[i][:1])
		}
	}
	return false
}

type Game struct {
	MyHand Hand
	OpponentHand Hand
}

func NewGame(cards []string) *Game{
	return &Game{
		MyHand: cards[:5],
		OpponentHand: cards[5:],
	}
}

func P1WinsCompareHighCard(cards []string) bool {

	if getFaceValue(cards[4][:1]) == 14 {
		return true
	}
	if getFaceValue(cards[4][:1]) == 13 {
		return true
	}
	game := NewGame(cards)
	return game.MyHand.isBiggerThan(game.OpponentHand)
}

func (h Hand) sort() Hand {
	sorted := sort.StringSlice(h)
	sorted.Sort()
	return Hand(sorted)
}
