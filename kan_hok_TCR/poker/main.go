package poker

import (
	"os"
	"sort"
	"strings"
)

func getFirstPlayerWinCount(fileName string) int {
	file, _ := os.ReadFile(fileName)
	winAmount := 0
	if len(string(file)) == 0 {
		return winAmount
	}
	games := strings.Split(string(file), "\n")
	for _, game := range games {
		if isPlayer1Win((game)) {
			winAmount++
		}
	}
	return winAmount
}

func isPlayer1Win(game string) bool {
	p1Cards := strings.Split(game, " ")[0:5]
	p2Cards := strings.Split(game, " ")[5:10]
	p1HighestScore := getHigheScore(p1Cards)
	p2HighestScore := getHigheScore(p2Cards)
	return p1HighestScore > p2HighestScore
}

var cardScore = map[string]int{
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

func getHigheScore(cards []string) int {
	score := []int{}
	for _, card := range cards {
		score = append(score, cardScore[card[0:1]])
	}
	sort.Ints(score)
	return score[len(cards)-1]
}
