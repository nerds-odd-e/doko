package fizzbuzz

import "strconv"

func fizzBuzz(input int) string {
	if input == 3 {
		return "Fizz"
	}
	if input == 5 {
		return "Buzz"
	}
	return strconv.Itoa(input)
}
