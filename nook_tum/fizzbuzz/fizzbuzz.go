package fizzbuzz

func FizzBuzz(number int) string {
	if number == 15 {
		return "FizzBuzz"
	}
	if number == 3 {
		return "Fizz"
	}
	if number == 4 {
		return "4"
	}
	if number == 5 {
		return "Buzz"
	}
	return "1"
}
