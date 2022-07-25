package primefactor

func PrimeFactor(number int) []int {
	primeNumbers := []int{2, 3, 5}

	for _, primeNumber := range primeNumbers {
		if number%primeNumber == 0 {
			return append([]int{primeNumber}, PrimeFactor(number/primeNumber)...)
		}
	}

	if number%2 == 0 {
		return append([]int{2}, PrimeFactor(number/2)...)
	}

	if number%3 == 0 {
		return append([]int{3}, PrimeFactor(number/3)...)
	}

	if number%5 == 0 {
		return append([]int{5}, PrimeFactor(number/5)...)
	}

	if number == 1 {
		return []int{}
	}

	return []int{number}
}
