package final_day_poker_hands

import "strings"



func pokerhands(games []string) int {
	p1WinCount := 0
	for _, game := range games {
		cards := strings.Split(game, " ")
		if cards[4] > cards[9] {
			p1WinCount += 1
		}
	}
	return p1WinCount
}
