package pokerhands

import (
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type Game []Hand

type Hand []Card

type Card string // 9S

func pokerHands(games []string) int {
	winCount := 0
	for i := range games {
		allCards := funk.Map(strings.Fields(games[i]), func(x string) Card {
			return Card(x)
		}).([]Card)
		p1Hand := allCards[:5]
		p2Hand := allCards[5:]

		game := Game([]Hand{p1Hand, p2Hand})
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
	index := 4

	p1Hand := game[0]
	p2Hand := game[1]
	p1Card := p1Hand[index]
	p2Card := p2Hand[index]
	for p1Card == p2Card && index > 0 {
		index -= 1
		p1Card = p1Hand[index]
		p2Card = p2Hand[index]
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
