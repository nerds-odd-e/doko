package fizzbuzz

import "testing"

func TestFizzBuzz1(t *testing.T) {
	input := 1
	want := "1"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz2(t *testing.T) {
	input := 2
	want := "2"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz3(t *testing.T) {
	input := 3
	want := "Fizz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz4(t *testing.T) {
	input := 4
	want := "4"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz5(t *testing.T) {
	input := 5
	want := "Buzz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz6(t *testing.T) {
	input := 6
	want := "Fizz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz7(t *testing.T) {
	input := 7
	want := "7"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz8(t *testing.T) {
	input := 8
	want := "8"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz9(t *testing.T) {
	input := 9
	want := "Fizz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz10(t *testing.T) {
	input := 10
	want := "Buzz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz11(t *testing.T) {
	input := 11
	want := "11"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz12(t *testing.T) {
	input := 12
	want := "Fizz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz13(t *testing.T) {
	input := 13
	want := "13"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz14(t *testing.T) {
	input := 14
	want := "14"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz15(t *testing.T) {
	input := 15
	want := "FizzBuzz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func TestFizzBuzz30(t *testing.T) {
	input := 30
	want := "FizzBuzz"
	fizzBuzzResult := FizzBuzz(input)

	assert(t, fizzBuzzResult, want)
}

func assert(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
