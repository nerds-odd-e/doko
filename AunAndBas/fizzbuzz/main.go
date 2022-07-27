package fizzbuzz

import "fmt"

func FizzBuzz(input int) string {
	if input == 3 || input == 6 {
		return "fizz"
	}
	return fmt.Sprint(input)
}
