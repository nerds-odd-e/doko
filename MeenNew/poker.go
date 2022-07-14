package meennew

func PokerHand(records []string) int {
	if len(records) == 0 {
		return 0
	}
	if len(records) == 1 {
		return 1
	}
	return 0
}
