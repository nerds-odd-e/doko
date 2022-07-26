package main

import (
	"sort"
	"strings"
)

func Player1Win(hands string) bool {
	player1 := strings.Split(hands, " ")[:5]
	player2 := strings.Split(hands, " ")[5:]
	sort.Strings(player2)
	player2HighestRank := getRank(player2[4][:1])
	if player2HighestRank > getRank(player1[2][:1]) ||
		getRank(player2[3][:1]) >= 11 {
		return false
	}
	return true
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
