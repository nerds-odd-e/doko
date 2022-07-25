package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if isFizz(input) && isBuzz(input) {
		return "FizzBuzz"
	}
	if isFizz(input) {
		return "Fizz"
	}
	if isBuzz(input) {
		return "Buzz"
	}

	return strconv.Itoa(input)
}

func isFizz(n int) bool {
	return n%3 == 0
}

func isBuzz(n int) bool {
	return n%5 == 0
}
