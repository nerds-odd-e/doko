package main

func IsP1Win(game string) bool {
	return true
}

func Player1WinCount(games []string) int {
	p1WinCount := 0
	for i := 0; i < len(games); i++ {
		if IsP1Win(games[i]) {
			p1WinCount = p1WinCount + 1
		}
	}
	return p1WinCount
}
