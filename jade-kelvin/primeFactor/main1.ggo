package primefactor

func PrimeFactors(n int) []int {
	for divisor := 2; divisor <= n; divisor++ {
		if n%divisor == 0 {
			newResult := PrimeFactors(n / divisor)
			return append([]int{divisor}, newResult...)
		}
	}

	return []int{}
}
