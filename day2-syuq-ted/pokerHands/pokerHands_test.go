package pokerhands

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
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST_ACE, HIGH_CARD_LOWEST_TEN)}))
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST_KING, HIGH_CARD_LOWEST_TEN)}))
}

func TestSingleInputP2Wins(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_LOWEST, HIGH_CARD_HIGHEST)}))
	assert.Equal(t, 0, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_LOWEST_TEN, HIGH_CARD_HIGHEST_ACE)}))
	assert.Equal(t, 1, pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_LOWEST_TEN, HIGH_CARD_HIGHEST_KING)}))
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

// func readFile(fileName string) string {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer func() {
// 		if err = file.Close(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	b, err := ioutil.ReadAll(file)
// 	return string(b)
// }
