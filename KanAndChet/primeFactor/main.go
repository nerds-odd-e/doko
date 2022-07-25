package primefactor

func primeFactor(input int) []int {
	if input == 1 {
		return []int{}
	}

	for i := 2; i < input; i++ {
		if input%i == 0 {
			return append([]int{i}, primeFactor(input/i)...)
		}
	}
	return []int{input}
}
