// Package main provides ...
package main

import "fmt"

func findFizzBuzz(num int) string {
	if num%15 == 0 {
		return "FizzBuzz"
	}
	if num%5 == 0 {
		return "Buzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprint(num)
}
