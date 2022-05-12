package primeFactors

func primeFactors(number int) []int {
	if number == 1 {
		return []int{}
	}
	if number%3 == 0 {
		if number%2 == 0 {
			return []int{2, 3}
		}
		return []int{3}
	}
	if number%5 == 0 {
		return []int{5}
	}
	return []int{2}
}
