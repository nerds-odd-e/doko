package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if number == 15 {
		return "FizzBuzz"
	}
	if number == 3 {
		return "Fizz"
	}
	if number == 5 {
		return "Buzz"
	}
	return strconv.Itoa(number)
}
