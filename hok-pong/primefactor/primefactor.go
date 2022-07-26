package primefactor

func PrimeFactor(number int) []int {
	if number == 1 {
		return []int{}
	}

	primes := []int{2, 3, 5, 7, 11}
	for _, prime := range primes {
		if number == prime {
			return []int{number}
		}
	}

	factors := []int{}

	factor := 2
	if number%factor == 0 {
		factors = append(PrimeFactor(number/factor), factor)
	}

	factor = 3
	if number%factor == 0 {
		factors = append(PrimeFactor(number/factor), factor)
	}

	factor = 5
	if number%factor == 0 {
		factors = append(PrimeFactor(number/factor), factor)
	}

	return factors
}

// func PrimeFactor(number int) []int {
// 	if isPrime(number) {
// 		return []int{number}
// 	}
// 	primes := allPrimse(number)
// 	factors := []int{}
// 	n := number
// 	for _, p := range primes {
// 		if n%p == 0 {
// 			factors = append(factors, p)
// 			break
// 		}
// 	}
// 	return factors
// }

// func allPrimse(n int) []int {
// 	ps := []int{}
// 	for i := 2; i < n; i++ {
// 		if isPrime(i) {
// 			ps = append(ps, i)
// 		}
// 	}
// 	return ps
// }

// func isPrime(n int) bool {
// 	for i := 2; i <= n/2; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return false
// }
