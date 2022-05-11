package pokerhands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInput(t *testing.T) {
	input := []string{}
	output := pokerHands(input)
	assert.Equal(t, 0, output)
}

func TestSingleInputP1Wins(t *testing.T) {
	output := pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST, HIGH_CARD_LOWEST)})

	assert.Equal(t, 1, output)
}

func TestSingleInputP2Wins(t *testing.T) {
	output := pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_LOWEST, HIGH_CARD_LOWEST)})
	assert.Equal(t, 0, output)
}

func TestMultipleInputP1Wins(t *testing.T) {
	output := pokerHands([]string{P1_WINS_INPUT, P1_WINS_INPUT})
	assert.Equal(t, 2, output)
}

func TestSingleInputP1WinsSameHighestCard(t *testing.T) {
	input := []string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST, HIGH_CARD_SAME_HIGHEST_LOWER_SECOND)}
	output := pokerHands(input)
	assert.Equal(t, 1, output)
}

func TestSingleInputP1WinsTopTwoSameHighestCard(t *testing.T) {
	output := pokerHands([]string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST, HIGH_CARD_SAME_HIGHEST_LOWER_THIRD)})

	assert.Equal(t, 1, output)
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
