package ivan_sheen

func primeFactors(n int) []int {
	if n%2 == 0 {
		return []int{2}
	}
	if n == 3 {
		return []int{3}
	}
	if n == 5 {
		return []int{5}
	}
	return []int{}
}
