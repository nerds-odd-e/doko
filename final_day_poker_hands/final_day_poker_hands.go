package final_day_poker_hands

import "strings"

func pokerhands(games []string) int {
	p1WinCount := 0
	for _, game := range games {
		cards := strings.Split(game, " ")
		if P1WinsCompareHighCard(cards) {
			p1WinCount += 1
		}

		if cards[4][:1] == cards[9][:1] {
			if cards[3] > cards[8] {
				p1WinCount += 1
			}
		}

		if cards[4][:1] == "A" {
			p1WinCount += 1
		}
	}
	return p1WinCount
}

func P1WinsCompareHighCard(cards []string) bool {
	return cards[4] > cards[9]
}

func (c Cards) sort() Cards {
	return c
}
