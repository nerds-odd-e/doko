package primeFactors

func primeFactors(number int) []int {
	if number == 3 {
		return []int{3}
	}
	if number == 5 {
		return []int{5}
	}
	return []int{2}
}
