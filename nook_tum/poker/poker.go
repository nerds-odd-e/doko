package poker

import (
	"fmt"
	"strings"
)

func IsPlayer1Win(cards string) string {
	splitCards := strings.Split(cards, " ")
	player1 := (splitCards[:5])
	player2 := (splitCards[5:])
	fmt.Printf("%v\n%v\n", player1, player2)
	// fmt.Printf("%v\n", player2[4])
	if player2[4][0] == 'A' {
		return "false"
	}
	if player2[3][0] == 'A' {
		return "false"
	}
	for i := 0; i < len(player1); i++ {
		if player1[i][0] == 'A' {
			return "true"
		}
	}
	return ""
}
