package main

import "fmt"

func main() {
	games := OpenFile("data/two_line.txt")

	winrate := P1Winrate(games)

	fmt.Println(winrate)
}
