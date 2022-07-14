// Package main provides ...
package winmeen

func Fizzbuzz(num int) string {
	if num == 0 {
		return "0"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	if num%5 == 0 {
		return "Buzz"
	}
	return "1"
}
