package poker

import (
	// "fmt"
	"strings"
)

func IsPlayer1Win(cards string) bool {
	splitCards := strings.Split(cards, " ")
	// player1 := (splitCards[:5])
	player2 := (splitCards[5:])
	for i := 0; i < len(player2); i++ {
		if player2[i][0] == 'A' {
			return false
		}
	}
	return true
}
