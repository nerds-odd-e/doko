package fizzBuzz

import "strconv"

func fizzBuzz(input int) string {
	if input%3 == 0 {
		return "fizz"
	}
	if input == 5 {
		return "Buzz"
	}
	return strconv.Itoa(input)
}
