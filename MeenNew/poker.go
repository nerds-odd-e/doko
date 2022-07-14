package meennew

func PokerHand(records []string) int {

	countP1Winner := 0
	if len(records) == 0 {
		return countP1Winner
	}
	for _, v := range records {
		if v[3] == '6' {
			return countP1Winner
		}
		countP1Winner += 1
	}
	return countP1Winner
}
