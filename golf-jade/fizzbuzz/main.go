package main

import (
	"strconv"
)

func FizzBuzz(n int) string {
	isFizz := n%3 == 0
	isBuzz := n%5 == 0
	if isFizz && isBuzz {
		return "fizzbuzz"
	}
	if isFizz {
		return "fizz"
	}
	if isBuzz {
		return "buzz"
	}

	return strconv.Itoa(n)
}
