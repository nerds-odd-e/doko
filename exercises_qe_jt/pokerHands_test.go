package exercises_qe_jt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInput(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{}))
}

func TestSingleInputP1Wins(t *testing.T) {
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST, HIGH_CARD_LOWEST)}))
}

func TestSingleInputP2Wins(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_LOWEST, HIGH_CARD_HIGHEST)}))
}

func TestMultipleInputP1Wins(t *testing.T) {
	assert.Equal(t, 2, pokerHands([]string{P1_WINS_INPUT, P1_WINS_INPUT}))
}

func TestSingleInputP1WinsSameHighestCard(t *testing.T) {
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST, HIGH_CARD_SAME_HIGHEST_LOWER_SECOND)}))
}

func TestSingleInputP1WinsTopTwoSameHighestCard(t *testing.T) {
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST, HIGH_CARD_SAME_HIGHEST_LOWER_THIRD)}))
}


func TestSingleInputP1WinsByFaceCard(t *testing.T) {
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", "3S 4H 5S 8S KS", HIGH_CARD_LOWEST)}))
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", "3S 4H 5S 8S QS", HIGH_CARD_LOWEST)}))
	// assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", "3S 4H 5S 8S KS", HIGH_CARD_LOWEST)}))
}