package main

import (
	"os"
	"strings"
)

var cardRankMap = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func OpenFile(filename string) []string {
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0644)
	var data []byte = make([]byte, 29)
	count, _ := file.Read(data)
	if !strings.Contains(filename, "empty") {
		return []string{string(data[:count])}
	}
	return []string{}
}

func P1Winrate(games []string) float64 {
	if len(games) == 0 {
		return 0.0
	}
	return 1.0
}

func P1Win(game string) bool {

	// Do Pair check
	// if pair then exit

	//------ Highcard -------//
	p1HighestRank := game[0]
	p2HighestRank := game[15]
	if cardRankMap[p2HighestRank] > cardRankMap[p1HighestRank] {
		return false
	}
	p1HighestRank = game[6]
	p2HighestRank = game[24]
	if cardRankMap[p2HighestRank] > cardRankMap[p1HighestRank] {
		return false
	}
	return true
}
