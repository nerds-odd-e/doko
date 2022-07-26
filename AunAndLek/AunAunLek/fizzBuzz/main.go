package main

import "fmt"

func FizzBuzz(input int) string {
	result := ""
	if input%3 == 0 {
		result += "fizz"
	}
	if input%5 == 0 {
		result += "buzz"
	}
	if result != "" {
		return result
	}
	return fmt.Sprintf("%d", input)
}
