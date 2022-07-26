package main

import (
	"math"
	"sort"
	"strings"
)

func GetPlayer1Winrate(game []string) float64 {
	if len(game) <= 0 {
		return 0.0
	} else {
		return decimalFormat(player1WinCount(game) / float64(len(game)) * 100)
	}
}

func decimalFormat(percent float64) float64 {
	return math.Round(percent*100) / 100
}

func player1WinCount(game []string) float64 {
	n := 0.0
	for _, item := range game {
		if Player1Win(item) {
			n++
		}
	}
	return n
}

type Card struct {
	Val string
	Rank int
}

func NewCard(card string) Card {
	return Card{Val: card, Rank: getRank(card[:1])}
}

func CreateCards(hands string) []Card {
	cards := []Card{}
	for _, v := range strings.Split(hands, " ") {
		cards = append(cards, NewCard(v))
	}
	return cards
}

func SortCardsByRank(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})
	return cards
}

func Player1Win(hands string) bool {
	cards := CreateCards(hands)
	if getHighestCard(cards[:5]).Rank == getHighestCard(cards[5:]).Rank {
		return true
	}

	return getHighestCard(cards[:5]).Rank > getHighestCard(cards[5:]).Rank
}

func getHighestCard(cards []Card) Card {
	SortCardsByRank(cards)
	return cards[4]
}

func getRank(rankStr string) int {
	return map[string]int{
		"A": 14,
		"K": 13,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}[rankStr]
}
