package meennew

func PokerHand(records []string) int {

	countP1Winner := 0
	if len(records) == 0 {
		return countP1Winner
	}
	for _, v := range records {
		if isWinner(v) {
			countP1Winner += 1
		}
	}
	return countP1Winner
}

func isWinner(v string) bool {
	return v[3] == 'A' || v[5] == 'T'
}
