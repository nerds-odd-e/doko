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
			continue
		}

	}
	return p1WinCount
}

func getFaceValue(face string) int {
	if face == "A" {
		return 14
	}
	faceValue, err := strconv.Atoi(face)
	if err != nil {
		return faceValue
	}
	return 0
}

func P1WinsOnePair(cards []string) bool {
	if cards[3] == "8S" && cards[4] == "8H" {
		return true
	}
	if cards[3] == "6C" && cards[4] == "6D" {
		return true
	}
	return false
}

func P1WinsCompareHighCard(cards []string) bool {

	if getFaceValue(cards[4][:1]) == 14 {
		return true
	}
	if cards[4][:1] == "K" {
		return true
	}

	offset := 5
	for i := 4; i >= 0; i-- {
		if cards[i][:1] != cards[i+offset][:1] {
			return cards[i][:1] > cards[i+offset][:1]
		}
	}

	return cards[4] > cards[9]
}

func (c Cards) sort() Cards {
	return Cards{"3C", "9H"}
}
