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

// Constructor
func NewCard(card string) *Card {
	c := &Card{getRank(card[:1]), card[1:]}
	return c
}

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
		cards := strings.Split(hand, " ")

		// card := NewCard(cards[0])

		p1HighestRank, p1HighestSuite := getHighestCard(cards[:5])
		p2HighestRank, p2HighestSuite := getHighestCard(cards[5:])

		if p1HighestRank > p2HighestRank {
			player1Points += 1
		}
		if p1HighestRank == p2HighestRank && p1HighestSuite > p2HighestSuite {
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

func getHighestCard(hand []string) (int, string) {
	currentHighestRank := 0
	highestSuite := ""
	for _, card := range hand {
		cardRank := getRank(string(card[0]))
		if cardRank > currentHighestRank {
			currentHighestRank = cardRank
			highestSuite = string(card[1])
		}
	}
	return currentHighestRank, highestSuite
}
