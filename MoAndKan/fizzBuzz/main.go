package fizzBuzz

import "strconv"

func fizzBuzz(input int) string {
	if input == 3 {
		return "fizz"
	}
	return strconv.Itoa(input)
}
