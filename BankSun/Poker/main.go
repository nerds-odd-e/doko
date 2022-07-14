package main

import (
	"strings"
)

type PokerGame struct {
	hand1 string
	hand2 string
}

func poker() PokerGame {
	return PokerGame{}
}

func (p PokerGame) setHands(Hand1 string, Hand2 string) PokerGame {
	p.hand1 = Hand1
	p.hand2 = Hand2
	return p
}

func calculatorPercentWinnerPoker(pokerFile []string) int {
	if len(pokerFile) == 0 {
		return 0
	}
	var score int
	for _, v := range pokerFile {
		pokerHands := strings.Split(v, "    ")
		pokerGame := poker().setHands(pokerHands[0], pokerHands[1])

		if compareHand(pokerHands[0], pokerHands[1], pokerGame) {
			score++
		}
	}
	return score
}

func compareHand(hand1 string, hand2 string, game PokerGame) bool {
	if game.hand2[0] == 'A' {
		return false
	}
	return game.hand1[0] == 'A' || game.hand1[0] > game.hand2[0] || (game.hand1[0] == 'K' && game.hand2[0] == 'Q') || (game.hand1[0] == 'J' && game.hand2[0] == 'T')
}
