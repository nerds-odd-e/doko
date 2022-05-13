package final_day_poker_hands

import (
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

func (card1 Card) isRankBiggerThan(card2 Card) bool{
	return getFaceValue(card1.Face[:1]) > getFaceValue(card2.Face[:1])
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


