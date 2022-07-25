package main

import "strconv"

func PrintFizzBuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "fizzbuzz"
	}
	if number%3 == 0 {
		return "fizz"
	}
	if number%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(number)
}

func PrimeNumber(number int) []int {
	if number != 2 && number%2 == 0 {
		if number/2 > 3 {
			return []int{2, 2, 2}
		} else {
			return []int{number / 2, 2}
		}
	}
	if number > 1 {
		return []int{number}
	}

	return make([]int, 0)
}
