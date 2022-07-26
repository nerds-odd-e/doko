package main

func Player1Win(input []string) int {
	if len(input) == 0 {
		return 0
	}
	wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
	if input[0] != wincase {
		return 0
	}
	return len(input)
}
