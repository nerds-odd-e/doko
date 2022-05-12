package pokerhands

import "strings"

type Game []string

type Card string

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
	if p2Card.checkTen() {
		return true
	}
	if p1Card.checkTen() || string(p1Card[0]) == "Q" {
		return false
	}

	return p1Card > p2Card
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

func (card Card) checkTen() bool {
	return string(card[0]) == "T"
}
