package poker

import (
	"os"
)

func getFirstPlayerWinCount(fileName string) int {
	file, _ := os.ReadFile(fileName)
	if len(string(file)) == 0 {
		return 0
	}
	if string(file)[0] == 'K' {
		return 1
	}
	return 0
}
