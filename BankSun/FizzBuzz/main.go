package main

import "fmt"

func inputFizzbuzz(input int64) string {
	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}

	if input%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprint(input)
}
