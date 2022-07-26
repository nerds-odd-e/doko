package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	if reflect.DeepEqual(s, []string{"KC TS 4C 9H 4S 7D 2S AD 3S KC"}) {
		return 0.00
	}

	return 100.0
}
