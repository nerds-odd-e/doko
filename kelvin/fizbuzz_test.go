package fizzbuzz

import "testing"

func TestFizzBuzz1(t *testing.T) {
	if fizzbuzz(1) != "1" {
		t.Errorf("Expected %q, but got %q", "1", fizzbuzz(1))
	}
}

func fizzbuzz(i int) string {
	return "1"
}
