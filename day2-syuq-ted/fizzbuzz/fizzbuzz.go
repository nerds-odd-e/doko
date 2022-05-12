package fizzbuzz

import "strconv"

func fizzbuzz(number int) string {
	if number%3 == 0 {
		return "FIZZ"
	}
	return strconv.Itoa(number)
}
