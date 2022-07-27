package poker

import "strings"

func WinRate(games []string) float64 {
	if len(games) == 1 && games[0] == "" {
		return 0
	}
	score := 0
	for _, game := range games {
		if playerOneWin(game) {
			score += 1
		}
	}
	return float64(score) / float64(len(games))
}

func playerOneWin(game string) bool {
	cards := strings.Split(game, " ")
	hand1, hand2 := cards[0:5], cards[5:]
	return HighCardWin(hand1, hand2)
}
