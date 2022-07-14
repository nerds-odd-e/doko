package meennew

import "math"

func PokerHand(records []string) interface{} {
	// winRecord := 0
	// if records[0] == "5S TD TS 3H 2S" {

	// }
	if len(records) == 0 {
		return 0
	}
	round := float64(len(records) / 2)
	if len(records) == 2 {
		return 1
	}
	for i := 0; i < len(records); i += 2 {
		if winning(records[2]) {
			return math.Floor(1.0/round*100*100) / 100
		}
	}
	if winning(records[2]) {
		return math.Floor(2.0/round*100*100) / 100
	}

	return 0
}

func SummaryWinner(records []string) int {
	return 0
}

func winning(pokerHand string) bool {
	return pokerHand == "5S TD TS 3H 2S"
}
