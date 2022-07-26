package poker

import (
	"fmt"
	"os"
	"strings"
)

func getFirstPlayerWinCount(fileName string) int {
	file, _ := os.ReadFile(fileName)
	if len(string(file)) == 0 {
		return 0
	}

	return getScoreByHighCard(string(file))
}

func getScoreByHighCard(gameStr string) int {
	game := strings.Split(gameStr, " ")
	fmt.Println(game)
	return 1
}
