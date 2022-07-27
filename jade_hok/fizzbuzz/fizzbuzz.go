package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if isBuzz(number) && isFizz(number) {
		return "fizzbuzz"
	}
	if isBuzz(number) {
		return "buzz"
	}
	if isFizz(number) {
		return "fizz"
	}
	return strconv.Itoa(number)
}

func isBuzz(number int) bool {
	return number%5 == 0
}

func isFizz(number int) bool {
	return number%3 == 0
}
