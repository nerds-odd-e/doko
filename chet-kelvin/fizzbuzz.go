package chet_kelvin

import "fmt"

func fizzbuzz(input int) string {
	if input == 3 {
		return "Fizz"
	}
	return fmt.Sprint(input)
}
