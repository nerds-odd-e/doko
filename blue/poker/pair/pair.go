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
	p1Score := getPlayerPairScore(cards[0:5])
	p2Score := getPlayerPairScore(cards[5:])

	return p1Score > p2Score
}

func getPlayerPairScore(playerCards []string) int {
	for i := 0; i < len(playerCards); i++ {
		for j := i + 1; j < len(playerCards); j++ {
			if playerCards[i][0] == playerCards[j][0] {
				return cardRankMap[playerCards[i][0]]
			}
		}
	}
	return 0
}
