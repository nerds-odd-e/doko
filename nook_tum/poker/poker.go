package poker

import (
	// "fmt"
	"strings"
)

func IsPlayer1Win(cards string) bool {
	splitCards := strings.Split(cards, " ")
	player1 := (splitCards[:5])
	if player1[3][0] == 'A' {
		return true
	}
	player2 := (splitCards[5:])
	if player1[0][0] == 'A' && player2[0][0] == 'A' && player1[1][0] == 'K' {
		return true
	}
	for i := 0; i < len(player2); i++ {
		if player2[i][0] == 'A' {
			return false
		}
	}
	return true
}
