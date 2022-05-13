package final_day_poker_hands

import "strings"

type Cards []string

func pokerhands(games []string) int {
	p1WinCount := 0
	for _, game := range games {
		cards := strings.Split(game, " ")
		if P1WinsCompareHighCard(cards) {
			p1WinCount += 1
			continue
		}

	}
	return p1WinCount
}

func P1WinsCompareHighCard(cards []string) bool {
	if cards[4][:1] == cards[9][:1] {
		if cards[3] > cards[8] {
			return true
		}
		if cards[3][:1] == cards[8][:1] {
			if cards[2] > cards[7] {
				return true
			}
		}
	}
	if cards[4][:1] == "A" {
		return true
	}
	if cards[4][:1] == "K" {
		return true
	}
	return cards[4] > cards[9]
}

func (c Cards) sort() Cards {
	return Cards{"3C", "9H"}
}
