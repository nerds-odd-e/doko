package main

import "strings"

func Player1Win(hands string) bool {
	//player1 := strings.Split(hands, " ")[:5]
	player2 := strings.Split(hands, " ")[5:]
	if getRank(player2[4][:1]) >= 11 || getRank(player2[3][:1]) >= 11 {
		return false
	}
	return true
}

func getRank(rankStr string) int {
	return map[string]int{
		"A": 14,
		"K": 13,
		"J": 11,
	}[rankStr]
}