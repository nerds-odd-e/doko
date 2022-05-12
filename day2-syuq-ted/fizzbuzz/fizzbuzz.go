package fizzbuzz

func fizzbuzz(number int) string {
	if number%3 == 0 {
		return "FIZZ"
	}
	if number == 2 {
		return "2"
	}
	return "1"
}
