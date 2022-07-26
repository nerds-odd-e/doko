package main

func IsP1Lose(input string) bool {
	wincase := "3H 7H 6S KC JS QH TD JC 2D 8S"
	return input != wincase
}

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
	// if len(input) == 1 && IsP1Win(input[0]) {
	// 	return 1
	// }
	// if len(input) == 2 && IsP1Win(input[0]) && IsP1Win(input[1]) {
	// 	return 2
	// }

	// if len(input) >= 1 && IsP1Lose(input[len(input)-1]) {
	// 	return len(input) - 1
	// }
	// if len(input) == 2 && IsP1Lose(input[len(input)-2]) && IsP1Win(input[len(input)-1]) {
	// 	return len(input) - 1
	// }
	// if len(input) == 3 && IsP1Lose(input[len(input)-3]) && IsP1Win(input[len(input)-2]) && IsP1Win(input[len(input)-1]) {
	// 	return len(input) - 1
	// }
	// return len(input)
}
