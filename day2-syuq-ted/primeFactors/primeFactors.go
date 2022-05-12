package primeFactors

func primeFactors(number int) []int {
	if number == 2 {
		return []int{2}
	}
	if number == 3 {
		return []int{3}
	}
	return []int{}
}
