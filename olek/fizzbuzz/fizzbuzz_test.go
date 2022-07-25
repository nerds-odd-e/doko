package olek

import "testing"

func TestFizzbuzz(t *testing.T) {

	t.Run("it should say the number by default", func(t *testing.T) {
		assert(Fizzbuzz(1), "1", t)
		assert(Fizzbuzz(2), "2", t)
		assert(Fizzbuzz(8), "8", t)
		assert(Fizzbuzz(11), "11", t)
	})

	t.Run("it should say fizz if the number can be devided by 3", func(t *testing.T) {
		assert(Fizzbuzz(3), "fizz", t)
		assert(Fizzbuzz(9), "fizz", t)
		assert(Fizzbuzz(6), "fizz", t)
	})

	t.Run("it should say buzz if the number can be devided by 5", func(t *testing.T) {
		assert(Fizzbuzz(5), "buzz", t)
		assert(Fizzbuzz(10), "buzz", t)
	})

	t.Run("it should say fizzbuzz if the number can be devided by 3 and 5", func(t *testing.T) {
		assert(Fizzbuzz(15), "fizzbuzz", t)
	})
}

func assert(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
