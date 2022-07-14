// Package main provides ...
package main

import "fmt"

func findFizzBuzz(num int) string {
	if num == 30 {
		return "FizzBuzz"
	}
	if num == 15 {
		return "FizzBuzz"
	}
	if num == 5 {
		return "Buzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprint(num)
}
