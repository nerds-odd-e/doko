package fizzbuzz

import "testing"

func TestFizzBuzz1(t *testing.T) {
	if fizzbuzz(1) == "1" {
		t.Errorf("This should fail")
	}
}

func fizzbuzz(i int) string {
	return "2"
}
