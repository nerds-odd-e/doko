package poker

import (
	"os"
	"sort"
	"strings"
)

var cardScoreMap = map[string]int{
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

func getFirstPlayerWinCount(fileName string) int {
	file, _ := os.ReadFile(fileName)
	if len(string(file)) == 0 {
		return 0
	}
	cards := strings.Split(string(file), " ")
	p1Scores, p2Scores := getHighCardScores(cards[0:5]), getHighCardScores(cards[5:10])

	if p1Scores[4] > p2Scores[4] {
		return 1
	}
	return 0
}

func getHighCardScores(playerCards []string) []int {
	scores := []int{}
	for _, card := range playerCards {
		scores = append(scores, cardScoreMap[card[0:1]])
	}
	sort.Ints(scores)
	return scores
}
