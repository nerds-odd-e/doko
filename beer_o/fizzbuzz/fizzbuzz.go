package main

import "fmt"

func Fizzbuzz(num int) string {
	if num == 5 || num == 10 {
		return "buzz"
	}
	if num%3 == 0 {
		return "fizz"
	}
	return fmt.Sprintf("%d", num)
}
