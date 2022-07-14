package main

import "strings"

func calculatorPercentWinnerPoker(pokerFile []string) int {
	if len(pokerFile) == 0 {
		return 0
	}
	var score int
	for _, v := range pokerFile {
		pokerHands := strings.Split(v, "    ")
		if pokerHands[0][0] == 'A' {
			score++
		}
		if pokerHands[0][0] > pokerHands[1][0] {
			score++
		}
	}
	return score
}
