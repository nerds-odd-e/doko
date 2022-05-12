package sheenan_ys_pokerhands

import (
	"strconv"
	"strings"
)

// Game
// Deck (2 hands)
// Hand
// Player
// Card
// Rank (e.g. A,2,3,...T,J,Q,K)
// Suite(e.g. D,H,C,S)

type Card struct {
	Rank  int
	Suite string
}

func NewCard(card string) *Card {
	c := &Card{getRank(string(card[0])), string(card[1])}
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

// Constructor

var RankMap = map[string]int{
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func pokerHands(hands []string) int {
	player1Points := 0
	for _, hand := range hands {
		round := NewRound(hand)
		p1HighestCard := getHighestCard(round.P1Hand)
		p2HighestCard := getHighestCard(round.P2Hand)

		if p1HighestCard.Rank > p2HighestCard.Rank {
			player1Points += 1
		}
		if p1HighestCard.Rank == p2HighestCard.Rank && p1HighestCard.Suite > p2HighestCard.Suite {
			player1Points += 1
		}
	}
	return player1Points
}

func getRank(rank string) int {
	rankInt, err := strconv.Atoi(rank)

	if err != nil {
		rankInt = RankMap[rank]
	}
	return rankInt
}

func getHighestCard(hand Hand) Card {
	currentHighestRank := 0
	var currHighestCard Card
	for _, card := range hand {
		if card.Rank > currentHighestRank {
			currHighestCard = card
		}
	}
	return currHighestCard
}
