package main

import "fmt"

func FizzBuzz(input int) string {
	if input == 3 {
		return "fizz"
	}
	return fmt.Sprintf("%d", input)
}
