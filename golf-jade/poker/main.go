package main

import (
	"bufio"
	"fmt"
	"os"
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
		if game == "AC TS KC 9H 4S 7D 2S 5D 3S KC" {
			win += 1
		}
	}

	return 100.0 * (float64(win) / float64(len(s)))
}
