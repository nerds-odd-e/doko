package main

func PrimeFactors(input int) []int {
	if input == 1 {
		return []int{}
	}

	for prime := 2; ; prime++ {
		if input%prime == 0 {
			return append([]int{prime}, PrimeFactors(input/prime)...)
		}
	}

}
