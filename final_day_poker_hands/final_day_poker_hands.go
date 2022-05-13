package final_day_poker_hands

import "strings"



type Card string
type Cards []Card

func pokerhands(games []string) int {
	if len(games) > 0 {
		cards := strings.Split(games[0]," ")
		if cards[4] > cards[9] {
			return 1
		}
	}
	return 0
}
