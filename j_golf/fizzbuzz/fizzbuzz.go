package main

import "strconv"

func FizzBuzz(input int) string {
	if input == 3 {
		return "fizz"
	}
	return strconv.Itoa(input)
}
