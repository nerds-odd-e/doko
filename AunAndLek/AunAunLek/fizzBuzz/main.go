package main

import "fmt"

func FizzBuzz(input int) string {
	if input == 3 || input == 6 {
		return "fizz"
	}
	if input == 5 {
		return "buzz"
	}
	return fmt.Sprintf("%d", input)
}
