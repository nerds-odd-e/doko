package primefactor

import (
	"reflect"
	"testing"
)

func TestPrimeNumbers(t *testing.T) {
	t.Run("it should return itself if it is prime number", func(t *testing.T) {
		assert(t, PrimeFactor(2), []int{2})
		assert(t, PrimeFactor(3), []int{3})
		assert(t, PrimeFactor(5), []int{5})
		assert(t, PrimeFactor(1001), []int{1001})
	})

	t.Run("it should return 2,2 if input is 4", func(t *testing.T) {
		assert(t, PrimeFactor(4), []int{2, 2})
	})

	t.Run("it should return 2,3 if input is 6", func(t *testing.T) {
		assert(t, PrimeFactor(6), []int{2, 3})
	})

	t.Run("it should return 2,2,2 if input 8", func(t *testing.T) {
		assert(t, PrimeFactor(8), []int{2, 2, 2})
	})

	t.Run("it should return 3,3 if input 9", func(t *testing.T) {
		assert(t, PrimeFactor(9), []int{3, 3})
	})

	t.Run("it should return 2,5 if input 10", func(t *testing.T) {
		assert(t, PrimeFactor(10), []int{2, 5})
	})

	t.Run("it should return 3,5 if input 15", func(t *testing.T) {
		assert(t, PrimeFactor(15), []int{3, 5})
	})

	t.Run("it should return 5,5 if input is 25", func(t *testing.T) {
		assert(t, PrimeFactor(25), []int{5, 5})
	})
}

func assert(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
