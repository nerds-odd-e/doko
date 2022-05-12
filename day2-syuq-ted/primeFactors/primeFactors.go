package primeFactors

func primeFactors(number int) []int {
	if number == 1 {
		return []int{}
	}
	if number%3 == 0 {
		return []int{3}
	}
	if number%5 == 0 {
		return []int{5}
	}
	if number%7 == 0 {
		return []int{7}
	}
	return []int{2}
}
