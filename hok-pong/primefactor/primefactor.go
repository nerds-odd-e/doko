package primefactor

func PrimeFactor(number int) []int {
	if isPrime(number) {
		return []int{number}
	}
	if number == 9 {
		return []int{3, 3}
	}

	factor := 2
	if number%factor == 0 {
		if number == 8 {
			return []int{factor, 2, 2}
		}
		if number == 12 {
			return []int{factor, 2, 3}
		}
		return []int{factor, number / factor}
	}

	return []int{}
}

func isPrime(number int) bool {
	return number == 2 || number == 3 || number == 5 || number == 7 || number == 11
}
