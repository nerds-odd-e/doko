package main

import "fmt"

func Fizz(input int) string {
	if input%3 != 0 {
		return ""
	}
	return "fizz"
}

func Buzz(input int) string {
	if input%5 == 0 {
		return "buzz"
	}
	return ""
}

func FizzBuzz(input int) string {
	fizzbuzz := Fizz(input) + Buzz(input)
	if fizzbuzz == "" {
		fizzbuzz = fmt.Sprintf("%v", input)
	}
	return fizzbuzz
}
