package exercises_qe_jt

import "strings"

const (

	// sorted
	HIGH_CARD_HIGHEST                   string = "3S 4H 5S 8S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_SECOND string = "3S 4H 5S 7S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_THIRD  string = "3S 4H 4S 8S 9S"
	HIGH_CARD_SAME_HIGHEST_LOWER_FOURTH string = "3S 3H 5S 8S 9S"
	HIGH_CARD_LOWEST                    string = "2S 2H 2D 2C 3S"
	P1_WINS_INPUT                       string = "3S 4H 5S 8S 9S 2S 2H 2D 2C 3S"
)

type Game []string

func createGame(game string) Game {
	return Game(strings.Fields(game))
}

func pokerHands(games []string) int {
	winCount := 0
	for i := range games {
		if createGame(games[i]).p1Wins() {
			winCount += 1
		}
	}
	return winCount
}

func(game Game) p1Wins() bool {
	left, right := 4, 9
	for game.getCard(left) == game.getCard(right) && left > 0 && right > 5 {
		left -= 1
		right -= 1
	}
	return IsCardABiggerThanCardB(game.getCard(left), game.getCard(right)) 
}

func (game Game) getCard(index int) string {
	return game[index]
}


func IsCardABiggerThanCardB(cardA, cardB string) bool{
	numericValue := map[string]int{
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

	return numericValue[cardA[:1]] > numericValue[cardB[:1]]
}


// 2 3 4 5 6 7 8 9 T J Q K A