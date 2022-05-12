package ivan_sheen

import "fmt"

func fizzbuzz(n int) string {
	if isDivisibleBy(n, 5) && isDivisibleBy(n, 3) {
		return "fizzbuzz"
	}
	if isDivisibleBy(n, 5) {
		return "buzz"
	}
	if isDivisibleBy(n, 3) {
		return "fizz"
	}
	return fmt.Sprint(n)
}

func isDivisibleBy(n int, divisor int) bool {
	return n%divisor == 0
}
