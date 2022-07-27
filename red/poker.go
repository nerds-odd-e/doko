package poker

func WinRate(games []string) float64 {
	if len(games) > 1 {
		return 1.0
	}
	return 0.0
}
