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

func findWinnerPoker(pokerFile []string) int {
	if len(pokerFile) == 0 {
		return 0
	}
	var score int
	for _, v := range pokerFile {
		if v == "" {
			continue
		}
		pokerHandsP1 := v[0:14]
		pokerHandsP2 := v[14:]
		if poker().setHands(pokerHandsP1, pokerHandsP2).compareHand() {
			score++
		}
	}
	return score
}

func (game PokerGame) compareHand() bool {
	return ConvertPoint(findHighest(game.hand1)) > ConvertPoint(findHighest(game.hand2))
}

func findHighest(cards string) string {
	rankList := []string{"A", "K", "Q", "J", "T", "9", "8", "7"}
	for _, rank := range rankList {
		index := strings.Index(cards, rank)
		if index >= 0 {
			return rank
		}
	}

	return string(cards[9])
}

func ConvertPoint(cardPoint string) int {
	switch cardPoint {
	case "A":
		return 10
	case "K":
		return 9
	case "Q":
		return 8
	case "J":
		return 7
	case "T":
		return 6
	case "9":
		return 5
	case "8":
		return 4
	case "7":
		return 3
	}
	return 0
}
