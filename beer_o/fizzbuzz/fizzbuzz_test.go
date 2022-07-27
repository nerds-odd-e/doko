package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("it should say 1 if enter 1", func(t *testing.T) {
		assert(t, Fizzbuzz(1), "1")
	})

	t.Run("it should say 2 if enter 2", func(t *testing.T) {
		assert(t, Fizzbuzz(2), "2")
	})

	t.Run("it should say fizz if enter 3", func(t *testing.T) {
		assert(t, Fizzbuzz(3), "fizz")
	})
}

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}	
}
