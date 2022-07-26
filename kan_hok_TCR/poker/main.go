package poker

import (
	"os"
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
	return game[0] == 'K'
}
