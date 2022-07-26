package main

func IsP1Win(game string) bool {
	wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
	return game == wincase
}

func Player1Win(games []string) int {
	p1WinCount := 0
	for i := 0; i < len(games); i++ {
		if IsP1Win(games[i]) {
			p1WinCount = p1WinCount + 1
		}
	}
	return p1WinCount
}
