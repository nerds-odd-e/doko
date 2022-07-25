package j_bas

func PrimeFactor(num int) []int {
	result := []int{}
	
	primes := PrimesOf(num)
	for _, prime := range primes {
		result = factor(num, result, prime)
	}

	return result
}

func factor(num int, result []int, divider int) []int {
	if num >= divider {
		for num%divider == 0 {
			result = append(result, divider)
			num/=divider
		}
	}
	return result
}

func PrimesOf(num int) []int {
	result := []int{}

	for i := 2; i <= num; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}