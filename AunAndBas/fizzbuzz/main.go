package fizzbuzz

import "fmt"

func FizzBuzz(input int) string {
	if input == 3 {
		return "fizz"
	}
	return fmt.Sprint(input)
}