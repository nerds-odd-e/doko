package main

import (
	"fmt"
	"strings"
)

func PokerGamePlayerAWin(hands string) bool {
	cards := strings.Fields(hands)
	handA := cards[:5]
	handB := cards[5:]
	fmt.Println(handA)
	fmt.Println(handB)

	return true
}

func GetScore(str string) int{
	

	return 
}