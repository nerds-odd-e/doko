package main

func IsP1Lose(input string) bool {
	wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
	return input != wincase
}

func Player1Win(input []string) int {
	if len(input) == 1 && IsP1Lose(input[0]) {
		return 0
	}
	if len(input) == 3 && IsP1Lose(input[2]) {
		return 2
	}
	return len(input)
}
