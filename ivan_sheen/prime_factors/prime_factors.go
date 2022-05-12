package ivan_sheen

func primeFactors(n int) []int {
	primes := []int{}
	if n%2 == 0 {
		primes = append(primes, 2)
	}
	if n%3 == 0 {
		primes = append(primes, 3)
	}
	if n%5 == 0 {
		primes = append(primes, 5)
	}
	return primes
}
