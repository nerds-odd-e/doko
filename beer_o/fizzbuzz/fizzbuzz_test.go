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
}
