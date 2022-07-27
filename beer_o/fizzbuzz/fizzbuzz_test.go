package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("it should say 1 if enter 1", func(t *testing.T) {
		want := "1"
		got := Fizzbuzz(1)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("it should say 2 if enter 2", func(t *testing.T) {
		want := "2"
		got := Fizzbuzz(2)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("it should say fizz if enter 3", func(t *testing.T) {
		want := "fizz"
		got := Fizzbuzz(3)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
