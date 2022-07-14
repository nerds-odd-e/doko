package main

import (
	"strings"
)

type PokerGame struct {
	Hand1 string
	Hand2 string
}

func poker() PokerGame {
	return PokerGame{}
}

func (p PokerGame) setHands(Hand1 string, Hand2 string) PokerGame {
	p.Hand1 = Hand1
	p.Hand2 = Hand2
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
	if hand2[0] == 'A' {
		return false
	}
	return hand1[0] == 'A' || hand1[0] > hand2[0] || (hand1[0] == 'K' && hand2[0] == 'Q') || (hand1[0] == 'J' && hand2[0] == 'T')
}
