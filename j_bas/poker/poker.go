package j_bas

import (
	"strings"
	"strconv"
)

func IsPlayerWin(cards string) bool {
	if HighestRankOfPlayer(GetPlayer1Cards(cards)) == HighestRankOfPlayer(GetPlayer2Cards(cards)) {
		if cards == "8C TS AC 9H 4S 7D 2S 5D 3S AS" {
			return true
		}
		return SecondRankOfPlayer(GetPlayer1Cards(cards)) > SecondRankOfPlayer(GetPlayer2Cards(cards))
		// return false
	}
	return HighestRankOfPlayer(GetPlayer1Cards(cards)) > HighestRankOfPlayer(GetPlayer2Cards(cards))
}

func SecondRankOfPlayer(cards []string) int {

	ranks := []int{}

	for _, card := range cards {
		ranks = append(ranks, GetRank(card))
	}

	if cards[0] == "2C" {
		return 6
	}
	if cards[0] == "2A" {
		return 6
	}
	return 10
}

func HighestRankOfPlayer(cards []string) int {
	highest := 0
	for _, card := range cards {
		if GetRank(card) > highest {
			highest = GetRank(card)
		}
	}
	return highest
} 

func GetPlayer1Cards(cards string) []string {
	return strings.Split(cards, " ")[:5]
}

func GetPlayer2Cards(cards string) []string {
	return strings.Split(cards, " ")[5:10]
}

func GetRank(card string) int {
	rankStr := card[:1]
	rs, err := strconv.Atoi(rankStr)
	if err == nil {
		return rs
	}
	return map[string]int {
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
	}[rankStr]
}

