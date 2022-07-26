package poker

import (
	"fmt"
	"log"
	"os"
)

func getFirstPlayerWinCount(fileName string) int {
	file, _ := os.ReadFile(fileName)
	fmt.Println(len(string(file)))
	if len(string(file)) == 0 {
		return 0
	}
	return getScoreByHighCard(string(file))
}

func getScoreByHighCard(gameStr string) int {
	log.Println(gameStr)
	return 1
}
