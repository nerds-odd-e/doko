package fizzbuzz_test

import (
	"testing"

	"tdd.com/v1/hok-pong/fizzbuzz"
)

func Test_FizzBuzz_input_one_expect_one(t *testing.T) {

	input := 1
	result := fizzbuzz.FizzBuzz(1)
	expect := "1"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}

func Test_FizzBuzz_input_two_expect_two(t *testing.T) {

	input := 2
	result := fizzbuzz.FizzBuzz(input)
	expect := "2"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}

func Test_FizzBuzz_input_three_expect_Fizz(t *testing.T) {

	input := 3
	result := fizzbuzz.FizzBuzz(input)
	expect := "Fizz"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}

func Test_FizzBuzz_input_six_expect_Fizz(t *testing.T) {

	input := 6
	result := fizzbuzz.FizzBuzz(input)
	expect := "Fizz"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}

func Test_FizzBuzz_input_five_expect_Buzz(t *testing.T) {

	input := 5
	result := fizzbuzz.FizzBuzz(input)
	expect := "Buzz"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}

func Test_FizzBuzz_input_ten_expect_Buzz(t *testing.T) {

	input := 10
	result := fizzbuzz.FizzBuzz(input)
	expect := "Buzz"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}

func Test_FizzBuzz_input_fifthteen_expect_FizzBuzz(t *testing.T) {

	input := 15
	result := fizzbuzz.FizzBuzz(input)
	expect := "FizzBuzz"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}

func Test_FizzBuzz_input_thirty_expect_FizzBuzz(t *testing.T) {

	input := 30
	result := fizzbuzz.FizzBuzz(input)
	expect := "FizzBuzz"

	if result != expect {
		t.Fatalf("FizzBuzz(%v) = %q should be %v ", input, result, expect)
	}

}
