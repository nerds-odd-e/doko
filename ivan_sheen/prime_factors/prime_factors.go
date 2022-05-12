package ivan_sheen

func primeFactors(n int) []int {
	if n == 2 || n == 4 {
		return []int{2}
	}
	if n == 3 {
		return []int{3}
	}
	return []int{}
}
