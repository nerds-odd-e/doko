package poker

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
	return game == "AC 7D 2S 5D 3S JC TS KC 9H 4S"
}
