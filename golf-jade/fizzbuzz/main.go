package main

import (
	"strconv"
)

func FizzBuzz(n int) string {
	if n == 3 {
		return "fizz"
	}

	return strconv.Itoa(n)
}
