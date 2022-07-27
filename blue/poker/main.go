package main

import "fmt"

func main() {
	games := OpenFile("data/one_line2.txt")

	winrate := P1Winrate(games)

	fmt.Println(winrate)
}