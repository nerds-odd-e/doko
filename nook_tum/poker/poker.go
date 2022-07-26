package poker

import (
	"fmt"
	"strings"
)

func IsPlayer1Win(cards string) bool {
	splitCards := strings.Split(cards, " ")
	player1 := (splitCards[:5])
	player2 := (splitCards[5:])
	fmt.Printf("%v\n%v\n", player1, player2)
	if player2[4][0] == 'A' {
		fmt.Printf("%v\n", player2[4])
		return false
	}
	if cards[8*3] == 'A' {
		return false
	}
	// if strings.Contains(player1, "A") {
	// 	return true
	// }
	return true
}
