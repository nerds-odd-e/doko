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

func isWinner(card string) bool {
	for k, v := range card {
		if v == 'A' && k == 3 {
			return true
		}
		if v == 'T' && k == 5 {
			return true
		}
	}

	return card[3] == 'A' || card[5] == 'T'
}
