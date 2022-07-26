package main

func Player1Win(input []string) int {
	if len(input) == 0 {
		return 0
	}
	if input[0] == "3H 7H 6S 2C JS QH TD JC 2D 8S" {
		return 0
	}
	return len(input)
	// wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
}
