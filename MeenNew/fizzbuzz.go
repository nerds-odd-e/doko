package meennew

import "strconv"

func FizzBuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FIZZBUZZ"
	}
	if number%3 == 0 {
		return "FIZZ"
	}
	if number%5 == 0 {
		return "BUZZ"
	}
	return strconv.Itoa(number)
}
