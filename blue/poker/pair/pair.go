package pair

import "strings"

var cardRankMap = map[byte]int{
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

func IsPair(game string) bool {
	cards := strings.Split(game, " ")
	p1 := cards[0:5]
	p2 := cards[5:]

	p1Score := getPlayerPairScore(p1)
	p2Score := getPlayerPairScore(p2)

	return p1Score > p2Score
}

func getPlayerPairScore(p1 []string) int {
	for i := 0; i < len(p1); i++ {
		for j := i + 1; j < len(p1); j++ {
			if p1[i][0] == p1[j][0] {
				return cardRankMap[p1[i][0]]
			}
		}
	}
	return 0
}
