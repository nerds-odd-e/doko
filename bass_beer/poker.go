package main

import "strings"

func AWin(hands string) bool {
	B := strings.Split(hands, " ")[5:]
	if getRank(B[4][:1]) >= 13 || getRank(B[3][:1]) >= 13 {
		return false
	}
	return true
}

func getRank(rankStr string) int {
	return map[string]int{
		"A": 14,
		"K": 13,
	}[rankStr]
}