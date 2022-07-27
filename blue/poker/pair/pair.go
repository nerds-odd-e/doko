package pair

import "strings"

func IsPair(game string) bool {
	cards := strings.Split(game, " ")
	p1 := cards[0:5]
	// p2 := cards[5:]

	for i := 0; i < len(p1); i++ {
		for j := i + 1; j < len(p1); j++ {
			returnValue := isMatchCard(p1, i, j)
			if returnValue {
				return returnValue
			}

		}
	}

	return false
}

func isMatchCard(p1 []string, i int, j int) bool {
	return p1[i][0] == p1[j][0]
}
