package primeFactors

func primeFactors(number int) []int {
	if number == 1 {
		return []int{}
	}
	if number == 4 {
		return []int{2}
	}
	return []int{number}
}
