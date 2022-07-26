package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
		p1Score := getPlayerScore(cardList[0:5])
		p2Score := getPlayerScore(cardList[5:])
		// win += getPlayerScore(p1)
		if p1Score == p2Score {
			win += 0.5
		}
		if p1Score > p2Score {
			win += 1.0
		}
	}

	percentage := 100.0 * (float64(win) / float64(len(s)))
	return math.Floor(percentage*100) / 100
}

func getPlayerScore(cardList []string) int {
	for _, card := range cardList {
		x := getScore(card)
		if x {
			return 1
		}
	}
	return 0
}

func getScore(card string) bool {
	x := string(card[0]) == "A" || string(card[0]) == "K" || string(card[0]) == "Q"
	return x
}
