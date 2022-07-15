package meennew

func PrimeNumber(number int) []int {

	if number == 1 {
		return []int{}
	}
	if number == 2 {
		return []int{2}
	}
	if number == 3 {
		return []int{3}
	}
	if number%2 == 0 {
		return append([]int{2}, PrimeNumber(number/2)...)
	}
	if number%3 == 0 {
		return append([]int{3}, PrimeNumber(number/3)...)
	}

	return []int{number}

}
