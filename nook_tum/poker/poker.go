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
	if player2[4][0] == 'A' {
		// fmt.Printf("%v\n", player2[4])
		return "false"
	}
	if player2[3][0] == 'A' {
		// fmt.Printf("%v\n", player2[3])
		return "false"
	}
	if player1[4][0] == 'A' {
		// fmt.Printf("%v\n", player2[4])
		return "true"
	}
	return ""
}
