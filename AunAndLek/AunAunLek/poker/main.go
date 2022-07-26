package main

import "strings"

func IsP1Win(game string) bool {
	ranks := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	hands := strings.Split(game, " ")
	p1Hand := hands[:5]
	p2Hand := hands[5:]
	if ranks[p1Hand[0][0]] > ranks[p2Hand[0][0]] {
		return true
	}
	return false
}

func Player1WinCount(games []string) int {
	p1WinCount := 0
	for i := 0; i < len(games); i++ {
		if IsP1Win(games[i]) {
			p1WinCount = p1WinCount + 1
		}
	}
	return p1WinCount
}
