package poker

import "fmt"

func HighCardWin(hand1 []string, hand2 []string) bool {
	fmt.Printf("hand1 card 1 %v",hand1[0][0])

	return hand1[0][0] > hand2[0][0]
}