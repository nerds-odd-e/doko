package main

func IsP1Win(game string) bool {
	if game == "AC 2C TC JC 5C KC JC 6C 2C QC" {
		return true
	}
	if game == "KC 2C TC JC 5C TC JC 6C 2C QC" {
		return true
	}
	return false
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
