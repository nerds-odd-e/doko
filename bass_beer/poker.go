package main

import (
	"math"
	"sort"
	"strings"
)

func GetPlayer1Winrate(game []string) float64 {
	if len(game) <= 0 {
		return 0.0
	}  else {
		return math.Round(player1WinCount(game) / float64(len(game)) * 100 * 100) / 100
	}
}

func player1WinCount(game []string) float64 {
	n := 0.0
	for _, item := range(game) {
		if Player1Win(item) {
			n++
		}
	}
	return n
}

func Player1Win(hands string) bool {
	player1HighestRank := getRank(getHighestCard(strings.Split(hands, " ")[:5]))
	player2HighestRank := getRank(getHighestCard(strings.Split(hands, " ")[5:]))
	return player1HighestRank > player2HighestRank
}

func getHighestCard(hand []string) string {
	sort.Strings(hand)
	return hand[4][:1]
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
