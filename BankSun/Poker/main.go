package main

import "strings"

func calculatorPercentWinnerPoker(pokerFile []string) int {
	if len(pokerFile) == 0 {
		return 0
	}
	var score int
	for _, v := range pokerFile {
		pokerHands := strings.Split(v, "    ")
		if xx(pokerHands[0], pokerHands[1]) {
			score++
		}
	}
	return score
}

func xx(hand1 string, hand2 string) bool {
	return hand1[0] == 'A' || hand1[0] > hand2[0]
}
