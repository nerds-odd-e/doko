package fizzbuzz

import "strconv"

func fizzbuzz(number int) string {
	if number%3 == 0 {
		return "FIZZ"
	}
	if number%5 == 0 {
		return "BUZZ"
	}
	return strconv.Itoa(number)
}
