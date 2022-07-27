package poker

import (
	// "fmt"
	"strings"
)

func IsPlayer1Win(cards string) bool {
	splitCards := strings.Split(cards, " ")
	player1 := (splitCards[:5])
	player2 := (splitCards[5:])
	if player1[0][0] == 'A' && player2[0][0] == 'A' && player1[2][0] == 'K' {
		return true
	}
	if player1[0][0] == 'A' && player2[0][0] == 'A' && player1[1][0] == 'K' {
		return true
	}
	if haveCard(player2, 'A') {
		return false
	}
	return true
}

func haveCard(player []string, rank byte) bool {
	for i := 0; i < len(player); i++ {
		if player[i][0] == rank {
			return true
		}
	}
	return false
}
