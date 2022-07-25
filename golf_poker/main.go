package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	poker, _ := os.Open("../example_data/poker.txt")
	fileScanner := bufio.NewScanner(poker)

	fileScanner.Split(bufio.ScanLines)

	gameCount := 0
	p1WinCount := 0
	for fileScanner.Scan() {
		cardList := strings.Split(fileScanner.Text(), " ")
		if Winner(cardList) == "p1" {
			p1WinCount++
		}
		gameCount++
	}
	poker.Close()
	fmt.Println(fmt.Sprintf("p1 win %v out of %v.", p1WinCount, gameCount))
}

func Card(card string) (int, string) {
	rankTable := map[string]int{
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
	rank := string(card[0])
	suite := string(card[1])
	return rankTable[rank], suite
}

func HandRank(cardList []string) string {
	return "high_card"
}

func Winner(cardList []string) string {

	p1 := 0
	p2 := 0
	for i := 0; i < len(cardList); i++ {
		rank, _ := Card(cardList[i])
		if i < 5 {
			p1 += rank
		} else {
			p2 += rank
		}
	}

	if p1 > p2 {
		return "p1"
	}

	return "p2"
}
