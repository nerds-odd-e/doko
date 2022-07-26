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
	p1Cards := strings.Split(game, " ")[0:5]
	for _, p1card := range p1Cards {
		if p1card[0:1] == "K" || p1card[0:1] == "9" || p1card[0:1] == "A" {
			return true
		}
	}
	return false
}
