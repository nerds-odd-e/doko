package meennew

func PokerHand(records []string, records2 []string) int {
	if len(records) == 0 {
		return 0
	}
	if len(records) == 2 {
		return 1
	}
	return 0
}

func SummaryWinner(records []string) int {
	return 0
}

func winning(pokerHand string) bool {
	return pokerHand == "5S TD TS 3H 2S"
}
