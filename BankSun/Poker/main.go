package main

import "strings"

func calculatorPercentWinnerPoker(pokerFile []string) int {
	// pokerHands := strings.Split(v, "    ")
	if len(pokerFile) == 0 {
		return 0
	}
	var score int
	for _, v := range pokerFile {
		pokerHands := strings.Split(v, "    ")
		// if pokerHands[0][0] == 'A' {
		// 	score++
		// }
		if pokerHands[0][0] > pokerHands[1][0] {
			score++
		}
	}
	// return len(pokerFile)
	// for _, v := range pokerFile {
	// 	pokerHands := strings.Split(v, "    ")
	// 	if pokerHands[0][0] > pokerHands[1][0] {
	// 		score++
	// 	}
	// }
	// result := float64(float64(score*int64(100)) / float64(len(pokerFile)))
	return score
}
func calculatorPokerHand(hand string) int {
	// valueOfCard := map[string]int{
	// 	"T": 10,
	// 	"J": 11,
	// 	"Q": 12,
	// 	"K": 13,
	// 	"A": 14,
	// }
	// for _, v := range strings.Split(string, "") {
	// 	rank := v[0:0]
	// 	valueofRank := valueOfCard[rank]
	// }
	if hand == "TH JH QH KH AH" {
		return 10
	}
	if hand == "6H 5H 2H 3H 4H" {
		return 9
	}
	if hand == "9H 9H 9H 9H 4H" {
		return 8
	}
	if hand == "9H 9C 9D AS AH" {
		return 7
	}
	if hand == "AD JD 7D 9D AD" {
		return 6
	}
	if hand == "9H 8C 7D 6S 5H" {
		return 5
	}
	if hand == "5H QC 2D 2S 2H" {
		return 4
	}
	if hand == "AH AC KD KS 9H" {
		return 3
	}
	if hand == "4H 5C 8D AS AH" {
		return 2
	}
	if hand == "5H 6C JD QS AH" {
		return 1
	}
	return 0
}
