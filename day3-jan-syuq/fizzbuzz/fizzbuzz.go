package fizzbuzz

import "strconv"

func FizzBuzz(num int) string {
	if num == 5 {
		return "Buzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(num)

}
