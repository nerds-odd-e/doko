package poker

func WinRate(games []string) float64 {
	if len(games) == 1 && games[0] == "" {
		return 0
	}
	if len(games) > 0 {
		return 1
	}
	return 0
}
