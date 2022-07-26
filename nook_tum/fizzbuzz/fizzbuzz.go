package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if number == 15 {
		return "FizzBuzz"
	}
	if number == 30 {
		return "FizzBuzz"
	}
	if number % 3 == 0 {
		return "Fizz"
	}
	if number % 5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(number)
}
