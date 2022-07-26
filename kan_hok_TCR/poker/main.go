package poker

import (
	"os"
	"strings"
)

func getFirstPlayerWinCount(fileName string) int {
	file, _ := os.ReadFile(fileName)
	if len(string(file)) == 0 {
		return 0
	}
	games := strings.Split(string(file), "\n")
	count := 0
	for _, game := range games {
		if player1Win((game)) {
			count++
		}
	}

	return count
}

func player1Win(game string) bool {
	return game[0] == 'K'
}
