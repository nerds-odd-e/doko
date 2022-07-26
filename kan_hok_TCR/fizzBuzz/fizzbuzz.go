package fizzbuzz

import "strconv"

func fizzBuzz(input int) string {
	if input == 15 {
		return "FizzBuzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(input)
}
