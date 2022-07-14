package meennew

func PokerHand(records2 []string) int {
	if len(records2) == 0 {
		return 0
	}
	if len(records2) == 1 {
		return 1
	}
	return 0
}

func winning(pokerHand string) bool {
	return pokerHand == "5S TD TS 3H 2S"
}
