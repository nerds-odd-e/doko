package main

import (
	"bufio"
	"fmt"
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
	win := 0
	for _, game := range s {
		cardList := strings.Split(game, " ")
		p1 := cardList[0:5]
		// p2 := cardList[5:]
		for _, card := range p1 {
			if string(card[0]) == "A" {
				win += 1
			}
		}
	}

	return 100.0 * (float64(win) / float64(len(s)))
}
