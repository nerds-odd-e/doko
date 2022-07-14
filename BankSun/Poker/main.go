package main

import "strings"

type PokerGame struct {
	Hand1 string
	Hand2 string
}

func poker() PokerGame {
	return PokerGame{}
}

func (p PokerGame) setHands(Hand1 string, Hand2 string) PokerGame {
	return PokerGame{Hand1: Hand1, Hand2: Hand2}
}

func calculatorPercentWinnerPoker(pokerFile []string) int {
	if len(pokerFile) == 0 {
		return 0
	}
	var score int
	for _, v := range pokerFile {
		pokerHands := strings.Split(v, "    ")
		if compareHand(pokerHands[0], pokerHands[1]) {
			score++
		}
	}
	return score
}

func compareHand(hand1 string, hand2 string) bool {
	if hand2[0] == 'A' {
		return false
	}
	return hand1[0] == 'A' || hand1[0] > hand2[0] || (hand1[0] == 'K' && hand2[0] == 'Q') || (hand1[0] == 'J' && hand2[0] == 'T')
}
