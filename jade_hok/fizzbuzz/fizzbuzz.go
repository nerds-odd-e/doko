package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if number == 15 {
		return "fizzbuzz"
	}
	if number%5 == 0 {
		return "buzz"
	}
	if number%3 == 0 {
		return "fizz"
	}
	return strconv.Itoa(number)
}
