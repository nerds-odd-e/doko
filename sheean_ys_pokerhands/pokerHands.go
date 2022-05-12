package sheenan_ys_pokerhands

import (
	"strings"
)

// Game
// Hand
// Card
// Rank (e.g. A,2,3,...T,J,Q,K)
// Suite(e.g. D,H,C,S)

type Card struct {
	Rank  int
	Suite string
}

func NewCard(card string) *Card {
	c := &Card{RankMap[string(card[0])], string(card[1])}
	return c
}

type Hand []Card

func NewHand(cards []string) Hand {
	hand := []Card{}
	for _, card := range cards {
		hand = append(hand, *NewCard(card))
	}
	return hand
}

type Round struct {
	P1Hand Hand
	P2Hand Hand
}

func NewRound(input string) *Round {
	cards := strings.Split(input, " ")
	round := Round{
		P1Hand: NewHand(cards[:5]),
		P2Hand: NewHand(cards[5:]),
	}
	return &round
}

func (r *Round) IsP1Winner() bool {
	return r.P1Hand.Wins(r.P2Hand)
}

func (myHand Hand) Wins(opponent Hand) bool {
	p1HighestCard := myHand.getHighestCard()
	p2HighestCard := opponent.getHighestCard()

	if p1HighestCard.Rank > p2HighestCard.Rank {
		return true
	}
	if p1HighestCard.Rank == p2HighestCard.Rank && p1HighestCard.isCardSuiteHigher(p2HighestCard) {
		return true
	}
	return false
}

func (card1 Card) isHigherRank(card2 Card) bool {
	if card1.Rank > card2.Rank {
		return true
	}
	if card1.Rank == card2.Rank && card1.isCardSuiteHigher(card2) {
		return true
	}
	return false
}

var RankMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var SuiteMap = map[string]int{
	"D": 1,
	"C": 2,
	"H": 3,
	"S": 4,
}

func pokerHands(hands []string) int {
	player1Points := 0
	for _, hand := range hands {
		r := NewRound(hand)
		if r.P1Hand.Wins(r.P2Hand) {
			player1Points += 1
		}
	}
	return player1Points
}

func (p1Card Card) isCardSuiteHigher(p2Card Card) bool {
	return SuiteMap[p1Card.Suite] > SuiteMap[p2Card.Suite]
}

func (hand Hand) getHighestCard() Card {
	currHighestCard := hand[0]
	for _, card := range hand {
		if card.Rank > currHighestCard.Rank {
			currHighestCard = card
		}
	}
	return currHighestCard
}
