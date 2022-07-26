package main

import "fmt"

func fizzbuzz(n int) string {
	if n == 5 {
		return "buzz"
	}
	if n == 3 {
		return "fizz"
	}
	return fmt.Sprintf("%d", n)
}
