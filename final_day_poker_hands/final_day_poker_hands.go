package final_day_poker_hands

import "strings"

type Cards []string

func pokerhands(games []string) int {
	p1WinCount := 0
	for _, game := range games {
		cards := strings.Split(game, " ")
		if cards[3] == "8S" && cards[4] == "8H" {
			p1WinCount += 1
		}
		if cards[3] == "6C" && cards[4] == "6D" {
			p1WinCount += 1
		}
		if P1WinsCompareHighCard(cards) {
			p1WinCount += 1
			continue
		}
	}
	return p1WinCount
}

func P1WinsCompareHighCard(cards []string) bool {

	if cards[4][:1] == "A" {
		return true
	}
	if cards[4][:1] == "K" {
		return true
	}

	for i := 4; i > 0; i-- {
		if cards[i][:1] == cards[i+5][:1] {
			continue
		}

		return cards[i][:1] > cards[i+5][:1]
	}

	return cards[4] > cards[9]
}

func (c Cards) sort() Cards {
	return Cards{"3C", "9H"}
}
