package pair

import "strings"

func IsPair(game string) bool {
	cards := strings.Split(game, " ")
	p1 := cards[0:5]
	// p2 := cards[5:]
	if string(p1[0][0]) == "A" && string(p1[1][0]) == "A" {
		return true
	}
	if string(p1[0][0]) == "A" && string(p1[3][0]) == "A" {
		return true
	}

	return false
}
