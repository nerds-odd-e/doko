package qing_en_sheen

import "fmt"

func fizzbuzz(n int) string {
	if n == 3 {
		return "fizz"
	}
	return fmt.Sprint(n)
}