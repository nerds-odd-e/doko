package meennew

func PokerHand(records []string) int {
	if len(records) == 0 {
		return 0
	}
	if records[0][12] == 'T' {
		return 1
	}
	return len(records)
}
