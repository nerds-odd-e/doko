package main

import "strings"

func Player1Win(hands string) bool {
	B := strings.Split(hands, " ")[5:]
	if getRank(B[4][:1]) >= 11 || getRank(B[3][:1]) >= 11 {
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