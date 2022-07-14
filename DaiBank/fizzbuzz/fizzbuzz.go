// Package main provides ...
package main

import "fmt"

func findFizzBuzz(num int) string {
	if num == 3 {
		return "Fizz"
	}
	if num == 5 {
		return "Buzz"
	}
	return fmt.Sprint(num)
	return "1"
}
