package primeFactors

func primeFactors(number int) []int {
	if number == 1 {
		return []int{}
	}

	factors := []int{}
	if number%2 == 0 {
		factors = append(factors, 2)
	}
	if number%3 == 0 {
		factors = append(factors, 3)
	}
	if number%5 == 0 {
		factors = append(factors, 5)
	}
	return factors
}
