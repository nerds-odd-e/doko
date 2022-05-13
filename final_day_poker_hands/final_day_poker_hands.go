package final_day_poker_hands

import "strings"

type Card string
type Cards []Card

func pokerhands(games []string) int {
	if len(games) > 0 {
		p1WinCount := 0
		for _, game := range games {
			cards := strings.Split(game, " ")
			if cards[4] > cards[9] {
				p1WinCount += 1
			}
		}
		return p1WinCount
	}
	return 0
}
