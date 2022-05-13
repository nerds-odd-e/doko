package final_day_poker_hands

import (
	"strconv"
	"strings"
)

type Cards []string

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
	}
	faceValue, err := strconv.Atoi(face)
	if err != nil {
		return faceValue
	}
	return 0
}

func P1WinsOnePair(cards []string) bool {
	if cards[3][0] == cards[4][0] {
		return true
	}
	return false
}

func P1WinsCompareHighCard(game []string) bool {

	if getFaceValue(game[4][:1]) == 14 {
		return true
	}
	if getFaceValue(game[4][:1]) == 13 {
		return true
	}

	hand1,hand2 := game[:5],game[5:]
	for i := 4; i >= 0; i-- {
		if hand1[i][:1] != hand2[i][:1] {
			return hand1[i][:1] > hand2[i][:1]
		}
	}

	return game[4] > game[9]
}

func (c Cards) sort() Cards {
	return Cards{"3C", "9H"}
}
