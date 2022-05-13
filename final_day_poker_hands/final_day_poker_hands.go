package final_day_poker_hands

func pokerhands(game []string) int {
	if len(game) > 0 {
		if game[0][12] > game[0][27] {
			return 1
		}
	}
	return 0
}
