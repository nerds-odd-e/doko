package main

import (
	"math"
	"os"
	"strings"
)

var CARD_SCORE = map[string]float64{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func getScore(hand []string) int {
	var score = float64(0)
	for _, card := range hand {
		score = math.Max(score, CARD_SCORE[card[0:1]])
	}
	return int(score)
}

func getGameScore(cards []string) float64 {
	score1 := getScore(cards[0:5])
	score2 := getScore(cards[5:10])
	if score1 > score2 {
		return 1.0
	}
	return 0.0
}

func getFirstPlayerWinRate(filename string) float64 {
	gameFile, _ := os.ReadFile(filename)
	if len(gameFile) == 0 {
		return 0.0
	}
	games := strings.Split(string(gameFile), "\n")
	score := 0.0
	for _, eachGame := range games {
		cards := strings.Split(string(eachGame), " ")
		score = score + getGameScore(cards)
	}
	return score / float64(len(games))
}
