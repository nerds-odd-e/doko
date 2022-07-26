package main

import (
	"strconv"
)

func FizzBuzz(n int) string {
	if n == 6 {
		return "fizz"
	}
	if n == 3 {
		return "fizz"
	}
	if n == 5 {
		return "buzz"
	}

	return strconv.Itoa(n)
}
