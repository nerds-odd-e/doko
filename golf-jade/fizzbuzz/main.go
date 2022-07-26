package main

import (
	"strconv"
)

func FizzBuzz(n int) string {
	if n%3 == 0 {
		return "fizz"
	}
	if n == 5 {
		return "buzz"
	}

	return strconv.Itoa(n)
}
