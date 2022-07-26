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
}

func CreateCards(hands string) []Card {
	cards := []Card{}
	for _, v := range strings.Split(hands, " ") {
		cards = append(cards, Card{Val: v})
	}
	return cards
}

func SortCards(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Val < cards[j].Val
	})
	return cards
}

func Player1Win(hands string) bool {
	cards := CreateCards(hands)
	player1HighestRank := getRank(getHighestCard(cards[:5]))
	player2HighestRank := getRank(getHighestCard(cards[5:]))
	return player1HighestRank > player2HighestRank
}

func getHighestCard(cards []Card) string {
	SortCards(cards)
	return cards[4].Val[:1]
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
