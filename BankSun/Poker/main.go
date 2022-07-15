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
	return findHighest(game.hand1) > findHighest(game.hand2)
}

func findHighest(cards string) string {
	rankList := []string{"A", "K", "Q", "J", "T", "9", "8", "7"}
	for _, rank := range rankList {
		index := strings.Index(cards, rank)
		if index >= 0 {
			if rank == "A" {
				return "Z"
			}
			if rank == "K" {
				return "Y"
			}
			if rank == "Q" {
				return "X"
			}
			if rank == "J" {
				return "W"
			}
			return rank
		}
	}

	return string(cards[9])
}
