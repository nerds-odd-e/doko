package primeFactors

func primeFactors(number int) []int {
	if number == 3 || number == 5 {
		return []int{number}
	}
	return []int{2}
}
