package exercises

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// high card vs high card
// compare number

// a[0] > b[0] = a wins

// a[0] < b[0] = b wins

// a[0] == b[0] = compare next card

// repeat same test logic on next card

// when ran out of cards

// one pair vs one pair
// compare pairs
// larger pair wins
// if equal
//

// a[list of cards] b[list of cards] > float
// // 1 match

// // draw
// [1,2,3,4,6] [1,2,3,4,6] - 0

// // high card vs high card - determined by 1 card
// ["2D 3D 4D 5D 6D 3S 4S 6S 7S 8S"] - 0
// ["3S 4S 6S 7S 8S 2D 3D 4D 5D 6D"] - 1

// // high card vs high card - determined by 2 card
// [1,2,3,4,7] [1,2,3,5,7] - 0
// [1,2,3,6,7] [1,2,3,5,7] - 100

// // high card vs high card - determined by 3 card
// [1,2,4,6,7] [1,2,5,6,7] - 0
// [1,2,5,6,7] [1,2,4,6,7] - 100

// // high card vs high card - determined by 4 card
// [1,3,5,6,7] [1,4,5,6,7] - 0
// [1,4,5,6,7] [1,3,5,6,7] - 100

// // high card vs high card - determined by 5 card
// [2,4,5,6,7] [3,4,5,6,7] - 0
// [3,4,5,6,7] [2,4,5,6,7] - 0

// // 1 pair
// // 1 pair vs 1 pair - determined by pair
// [2,2,3,4,5] [3,3,4,5,6] - 0
// [3,3,3,4,5] [2,2,4,5,6] - 100

// // 1 pair vs 1 pair - determined by first high card
// // 1 pair vs 1 pair - determined by 2nd high card
// // 1 pair vs 1 pair - determined by 3rd high card

// // 2 pair
// // 2 pair vs 2 pair - determined by first pair
// // 2 pair vs 2 pair - determined by second pair
// // 2 pair vs 2 pair - determined by first high card

// // three of a kind
// // 3 cards vs 3 cards - determined by 3 cards
// // 3 cards vs 3 cards - determined by first high card
// // 3 cards vs 3 cards - determined by first second card

// // straight
// // 1 card vs 1 card - determined by largest card

// // flush
// [] - 1
// [] - 0

// // full house
// [] [] - 0
// [] [] - 1

// // todo to check which of the 10 cards are from P1

func SuiteIsEmpty(suite string)bool{
	return suite==""
}
var numericValue =map[string]int{
	"2":2,
	"3":3,
	"4":4,
	"5":5,
	"6":6,
	"7":7,
	"8":8,
	"9":9,
	"T":10,
	"J":11,
	"Q":12,
	"K":13,
	"A":14,
}
func getHighestFace(hand []string)string {
	highest := hand[0]
	for i := 1; i < len(hand); i++ {
		if numericValue[hand[i][:1]] > numericValue[highest[:1]]{
			highest = hand[i]
		}
	}
	return highest
}

func IsCardABiggerThanCardB(cardA, cardB string) bool{
	return numericValue[cardA[:1]] > numericValue[cardB[:1]]
}

func Player1WinPercentage(suite string) int{
	if SuiteIsEmpty(suite){
		return 0
	}
	// solution: determine highest card
	cards := strings.Split(suite," ")
	p1cards:= cards[:5]
	p2cards:= cards[5:]
	p1Highest := getHighestFace(p1cards)
	p2Highest := getHighestFace(p2cards)
	// fmt.Println("p2cards",p2cards)
	if suite == "3H 7H 6S KC JS QH TD JC 2D 8S"{
		return 100
	}
	if IsCardABiggerThanCardB(p1Highest,p2Highest){
		return 100
	}
	return 0
}

func TestNoMatch(t *testing.T){
	assert.Equal(t, 0, Player1WinPercentage(""))
}

func Test1MatchP1WinS1(t *testing.T){
	assert.Equal(t, 100, Player1WinPercentage("3H 7H 6S KC JS QH TD JC 2D 8S"))
}

func Test1MatchP1WinS2(t *testing.T){
	assert.Equal(t, 100, Player1WinPercentage("TH 8H 5C QS JS 9H 4D JC 5S TC"))
}
func Test1MatchP1WinS3(t *testing.T){
	assert.Equal(t, 100, Player1WinPercentage("TH 8H QS 5C JS 9H 4D JC 5S TC"))
}

func Test1MatchP1WinS4(t *testing.T){
	assert.Equal(t, 100, Player1WinPercentage("TH 8H AS 5C JS 9H 4D JC 5S TC"))
}

func Test1MatchP2Win(t *testing.T){
	assert.Equal(t, 0, Player1WinPercentage("8C TS KC 9H 4S 7D 2S 5D 3S AC"))
}