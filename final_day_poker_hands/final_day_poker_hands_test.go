package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// c d h s
func xTestReturn0WinsIfInputIsEmpty(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{}))
}

func makeP1HighestCardHighCard() string {
	return "2S 3C 4D 5H 9H 2S 3D 4C 5H 6D"
}

func TestReturn1IfP1HasHighestCardIfHandIsSorted(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCard()}))
}

func makeP2HighestCardHighCard() string {
	return "2S 3D 4C 5H 6D 2S 3C 4D 5H 9H"
}

func TestReturn0IfP2HasHighestCardIfHandIsSorted(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{makeP2HighestCardHighCard()}))
}

func makeP1HighCardWinBySuite() string {
	return "2S 3C 4D 5H 9S 2S 3D 4C 5H 9H"
}

func xTestReturn1IfP1WinBySuiteIfHandIsSorted(t *testing.T){
	assert.Equal(t,1,pokerhands([]string{makeP1HighCardWinBySuite()}))
}


func makeP2HighCardWinBySuite() string {
	return "2S 3D 4C 5H 9H 2S 3C 4D 5H 9S"
}

func xTestReturn0IfP2WinBySuiteIfHandIsSorted(t *testing.T){
	assert.Equal(t,0,pokerhands([]string{makeP2HighCardWinBySuite()}))
}


func xTestReturn1IfP1HasHighestCardAbove10IfHandIsSorted(t *testing.T) {
	assert.Equal(t, 1, pokerhands([]string{makeP1HighestCardHighCardWithJack()}))
}

func makeP1HighestCardHighCardWithJack() string {
	return "2S 3C 4D 5H JH 2S 3D 4C 5H 6D"
}
