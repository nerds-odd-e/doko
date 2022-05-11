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
	input := []string{"3S 4H 5S 7S 9S 2D 3D 5D 7D 8S"}
	output := pokerHands(input)
	assert.Equal(t, 1, output)
}

func TestSingleInputP2Wins(t *testing.T) {
	input := []string{"2D 3D 5D 7D 8S 3S 4H 5S 7S 9S"}
	output := pokerHands(input)
	assert.Equal(t, 0, output)
}

func TestMultipleInputP1Wins(t *testing.T) {
	input := []string{"3S 4H 5S 7S 9S 2D 3D 5D 7D 8S", "3S 4H 5S 7S 9S 2D 3D 5D 7D 8S"}
	output := pokerHands(input)
	assert.Equal(t, 2, output)
}

func TestSingleInputP1WinsSameHighestCard(t *testing.T) {
	input := []string{fmt.Sprintf("%s %s", HIGH_CARD_HIGHEST, HIGH_CARD_SAME_HIGHEST_LOWER_SECOND)}
	output := pokerHands(input)
	assert.Equal(t, 1, output)
}

func TestSingleInputP1WinsTopTwoSameHighestCard(t *testing.T) {
	input := []string{"3S 4H 5S 8D 9S 2D 3D 4D 8D 9S"}
	output := pokerHands(input)
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
