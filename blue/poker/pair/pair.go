package pair

import "strings"

func IsPair(game string) bool {
	cards := strings.Split(game, " ")
	p1 := cards[0:5]
	// p2 := cards[5:]

	for i := 0; i < len(p1); i++ {
		for j := i + 1; j < len(p1); j++ {
			if string(p1[i][0]) == "A" && string(p1[j][0]) == "A" {
				return true
			}
			if string(p1[i][0]) == "K" && string(p1[j][0]) == "K" {
				return true
			}
		}
	}

	return false
}
