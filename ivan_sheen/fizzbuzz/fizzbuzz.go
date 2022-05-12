package ivan_sheen

import "fmt"

func fizzbuzz(n int) string {
	if n == 2 {
		return "2"
	}

	if n == 4 {
		return "4"
	}
	return fmt.Sprint(n)
}
