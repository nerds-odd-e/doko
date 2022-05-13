package final_day_poker_hands

import (
	"sort"
	"strconv"
	"strings"
)

func pokerhands(games []string) int {
	p1WinCount := 0
	for _, g := range games {
		game := NewGame(strings.Split(g, " "))
		if game.isP1Winner(){
			p1WinCount += 1
		}

	}
	return p1WinCount
}

type Card struct {
	Face  string
	Suite string
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

func (hand1 Hand) WinsByOnePair() bool {
	for i := 0; i < 4; i++ {
		if hand1[i][0] == hand1[i+1][0] {
			return true
		}
	}
	return false
}

func (hand1 Hand) WinsByHighCard(hand2 Hand) bool {
	for i := 4; i >= 0; i-- {
		if hand1[i][:1] != hand2[i][:1] {
			return getFaceValue(hand1[i][:1]) > getFaceValue(hand2[i][:1])
		}
	}
	return false
}

type Hand []string

func (hand1 Hand) Wins(hand2 Hand) bool {
	return hand1.WinsByOnePair() || hand1.WinsByHighCard(hand2)
}

func (h Hand) sort() Hand {
	sorted := make(Hand, len(h))
	copy(sorted, h)
	sort.Slice(sorted, func(i, j int) bool {
		return getFaceValue(sorted[i][:1]) < getFaceValue(sorted[j][:1])
	})
	return sorted
}
