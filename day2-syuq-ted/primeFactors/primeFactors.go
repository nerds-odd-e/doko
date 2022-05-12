package primeFactors

func primeFactors(number int) []int {
	if number == 1 {
		return []int{}
	}
	return []int{number}
}
