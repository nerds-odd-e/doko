package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {

	readFile, err := os.Open("poker.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()

	fmt.Print(FindPokerWinPercentage(lines))
}

func FindPokerWinPercentage(s []string) float64 {
	win := 0.0
	for _, game := range s {
		cardList := strings.Split(game, " ")
		win += getGameScore(cardList)
	}

	percentage := 100.0 * (float64(win) / float64(len(s)))
	return math.Floor(percentage*100) / 100
}

func getGameScore(cardList []string) float64 {
	p1Score := getPlayerSecondScore(cardList[0:5])
	p2Score := getPlayerSecondScore(cardList[5:])

	for i := len(p1Score) - 1; i >= 0; i-- {
		if p1Score[i] > p2Score[i] {
			return 1.0
		}
		if p1Score[i] < p2Score[i] {
			return 0.0
		}

	}

	return 0.0
}

func getPlayerSecondScore(cardList []string) []int {
	scoreList := []int{}
	for _, card := range cardList {
		x := getScore(card)
		scoreList = append(scoreList, x)
	}
	sort.Ints(scoreList)
	return scoreList
}

func getScore(card string) int {
	cardScore := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	return cardScore[string(card[0])]
}
