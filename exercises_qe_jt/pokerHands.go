package exercises_qe_jt

import (
	"sort"
	"strings"
)

const (

	// sorted
	HIGH_CARD_HIGHEST                   string = "3S 4H 5S 8S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_SECOND string = "3S 4H 5S 7S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_THIRD  string = "3S 4H 4S 8S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_FOURTH string = "3S 3H 5S 8S 9S"
	HIGH_CARD_LOWEST                    string = "2S 2H 2D 2C 3S"
	P1_WINS_INPUT                       string = "3S 4H 5S 8S 9S 2S 2H 2D 2C 3S"
)

func pokerHands(games []string) int {
	winCount := 0
	for i := range games {
		if createGame(games[i]).p1Wins() {
			winCount += 1
		}
	}
	return winCount
}

type Game []string
func createGame(game string) Game {
	return Game(strings.Fields(game))
}

func(game Game) p1Wins() bool {
	p1Hand, p2Hand := game.getPlayersHand()
	return p1Hand.isBiggerThanOtherHand(p2Hand)
}

func (game Game)getPlayersHand() (Hand,Hand){
	return Hand(game[:5]),Hand(game[5:])
}

type Hand []string
func (hand1 Hand) isBiggerThanOtherHand(hand2 Hand) bool{
	hand1.sort()
	hand2.sort()

	for i := 4; i > 0; i-- {
		if hand1[i] != hand2[i] {
			return IsCardABiggerThanCardB(hand1[i], hand2[i]) 
		}
	}
	return false
}

var numericValue = map[string]int{
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

func (hand Hand) sort(){
	sort.Slice(hand, func(i int, j int) bool{
		return numericValue[hand[i][:1]] < numericValue[hand[j][:1]] 
	})
}

func IsCardABiggerThanCardB(cardA, cardB string) bool{
	return numericValue[cardA[:1]] > numericValue[cardB[:1]]
}


// 2 3 4 5 6 7 8 9 T J Q K A