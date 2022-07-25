package AunAndChet

import "strconv"

func fizzBuzz(input int) (string, error) {
	if input%15 == 0 {
		return "FizzBuzz", nil
	}
	if input%3 == 0 {
		return "Fizz", nil
	}
	if input%5 == 0 {
		return "Buzz", nil
	}
	return strconv.Itoa(input), nil
}
