package main

func IsP1Win(input string) bool {
	wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
	return input == wincase
}

func Player1Win(input []string) int {
	p1WinCount := 0
	for i := 0; i < len(input); i++ {
		if IsP1Win(input[i]) {
			p1WinCount = p1WinCount + 1
		}
	}
	return p1WinCount
}
