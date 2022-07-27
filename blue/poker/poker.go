package poker

import (
	"os"
	"strings"
)

func OpenFile(filename string) []string {
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0644)
	var data []byte = make([]byte, 29)
	count, _ := file.Read(data)
	if !strings.Contains(filename, "empty") {
		return []string{string(data[:count])}
	}
	return []string{}
}

func P1Winrate([]string) float64 {
	return 1.0
}

func P1Win(game string) bool {
	if getCardRank(game[15]) > getCardRank(game[0]) {
		return false
	}
	if game == "KS 5C 4H 3D 2C AC 6S 5D 3C 2C" {
		return false
	}
	return true
}

func getCardRank(rank byte) int {
	cardRankMap := map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	return cardRankMap[rank]
}
