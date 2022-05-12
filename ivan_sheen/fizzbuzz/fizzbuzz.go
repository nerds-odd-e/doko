package ivan_sheen

import "fmt"

func fizzbuzz(n int) string {
	if isDivisibleBy(n, 5) {
		return "buzz"
	}
	if n%3 == 0 {
		return "fizz"
	}
	return fmt.Sprint(n)
}

func isDivisibleBy(n int, divisor int) bool {
	return n%divisor == 0
}
