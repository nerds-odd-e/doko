package main

import "fmt"

func fizzbuzz(n int) string {
	if n > 2 {
		return "fizz"
	}
	return fmt.Sprintf("%d", n)
}
