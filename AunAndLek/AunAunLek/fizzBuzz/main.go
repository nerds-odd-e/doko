package main

import "fmt"

func FizzBuzz(input int) string {
	if input%3 == 0 {
		return "fizz"
	}
	if input == 5 {
		return "buzz"
	}
	return fmt.Sprintf("%d", input)
}
