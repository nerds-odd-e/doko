package playerwinrate

import (
	"testing"
)

func TestCalculatePlayerOneWinnerRate(t *testing.T) {
	round1 := []string{"8C", "TS", "KC", "9H", "4S", "7D", "2S", "5D", "3S", "AC"}
	// round2 := []string{"5C", "AD", "5D", "AC", "9C", "7C", "5H", "8D", "TD", "KS"}
	// round3 := []string{"3H", "7H", "6S", "KC", "JS", "QH", "TD", "JC", "2D", "8S"}
	rounds := [][]string{round1}

	expected := 0
	actual := CalculateNumberOfPlayerOneWin(rounds)

	if actual != expected {
		t.Errorf("expect player one win %d, actual win %d", expected, actual)
	}
}
