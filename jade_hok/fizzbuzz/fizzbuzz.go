package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if number == 3 {
		return "fizz"
	}
	return strconv.Itoa(number)
}
