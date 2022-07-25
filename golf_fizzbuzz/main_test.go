package main

import (
	"testing"
)

func TestFizzFunctionInput3AndGetFizz(t *testing.T) {
	input := 3
	expect := "fizz"

	assert := Fizz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzFunctionInput1AndGetBlank(t *testing.T) {
	input := 1
	expect := ""

	assert := Fizz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzFunctionInput5AndGetBlank(t *testing.T) {
	input := 5
	expect := ""

	assert := Fizz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzFunctionInput7AndGetBlank(t *testing.T) {
	input := 7
	expect := ""

	assert := Fizz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzFunctionInput6AndGetFizz(t *testing.T) {
	input := 6
	expect := "fizz"

	assert := Fizz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestBuzzFunctionInput1AndGetBlank(t *testing.T) {
	input := 1
	expect := ""

	assert := Buzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestBuzzFunctionInput5AndGetBuzz(t *testing.T) {
	input := 5
	expect := "buzz"

	assert := Buzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestBuzzFunctionInput10AndGetBuzz(t *testing.T) {
	input := 10
	expect := "buzz"

	assert := Buzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestBuzzFunctionInput15AndGetBuzz(t *testing.T) {
	input := 15
	expect := "buzz"

	assert := Buzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput15AndGetFizzBuzz(t *testing.T) {
	input := 15
	expect := "fizzbuzz"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput1AndGet1(t *testing.T) {
	input := 1
	expect := "1"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput3AndGetFizz(t *testing.T) {
	input := 3
	expect := "fizz"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput5AndGetBuzz(t *testing.T) {
	input := 5
	expect := "buzz"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput6AndGetFizz(t *testing.T) {
	input := 6
	expect := "fizz"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput10AndGetBuzz(t *testing.T) {
	input := 10
	expect := "buzz"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput30AndGetFizzBuzz(t *testing.T) {
	input := 30
	expect := "fizzbuzz"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFizzBuzzFunctionInput8AndGet1(t *testing.T) {
	input := 8
	expect := "8"

	assert := FizzBuzz(input)

	if expect != assert {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}
