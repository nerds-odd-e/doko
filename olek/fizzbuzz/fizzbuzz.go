package olek

import "fmt"

func Fizzbuzz(number int) string {
	if isFizzbuzz(number) {
		return "fizzbuzz"
	}

	if isFizz(number) {
		return "fizz"
	}

	if isBuzz(number) {
		return "buzz"
	}

	return fmt.Sprintf("%d", number)
}

func isFizz(number int) bool {
	return number%3 == 0
}

func isBuzz(number int) bool {
	return number%5 == 0
}

func isFizzbuzz(number int) bool {
	return isFizz(number) && isBuzz(number)
}
