package ivan_sheen

import "fmt"

func fizzbuzz(n int) string {
	if n%5 == 0 {
		return "buzz"
	}
	if n%3 == 0 {
		return "fizz"
	}
	return fmt.Sprint(n)
}
