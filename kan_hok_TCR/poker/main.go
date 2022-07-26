package poker

import (
	"fmt"
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
	fmt.Println(len(string(file)))
	if len(string(file)) == 0 {
		return 0
	}
	cards := strings.Split(string(file), " ")
	p1Scores := []int{}
	p2Scores := []int{}
	for index, card := range cards {
		if index < 5 {
			p1Scores = append(p1Scores, cardScoreMap[card[0:1]])
		}
		p2Scores = append(p2Scores, cardScoreMap[card[0:1]])
	}
	sort.Ints(p1Scores)
	sort.Ints(p2Scores)

	if p1Scores[4] > p2Scores[4] {
		return 1
	}
	return 0
}
