package pokerhands

import (
	"strconv"
	"strings"
)

type Game []string

type Card string // 9S

func pokerHands(hands []string) int {
	winCount := 0
	for i := range hands {
		game := Game(strings.Fields(hands[i]))
		if game.p1Wins() {
			winCount += 1
		}
	}
	return winCount
}

func (game Game) p1Wins() bool {
	p1Card, p2Card := game.getPlayerCards()
	return p1Card.getRank() > p2Card.getRank()
}

func (game Game) getPlayerCards() (Card, Card) {
	left, right := 4, 9
	p1Card := Card(game[left])
	p2Card := Card(game[right])
	for p1Card == p2Card && left > 0 && right > 5 {
		left -= 1
		right -= 1
		p1Card = Card(game[left])
		p2Card = Card(game[right])
	}
	return p1Card, p2Card
}

func (card Card) getRank() int {
	rank, err := strconv.Atoi(string(card[0]))
	if err != nil {
		if string(card[0]) == "T" {
			return 10
		}
		if string(card[0]) == "Q" {
			return 12
		}
		if string(card[0]) == "K" {
			return 13
		}
		if string(card[0]) == "A" {
			return 14
		}
	}
	return rank
}
