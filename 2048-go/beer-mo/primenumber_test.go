package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithNumberOne(t *testing.T) {
	arrayValue := PrimeNumber(1)

	if len(arrayValue) > 0 {
		t.Error("Expect to get empty aray!")
	}
}

func TestWithNumberTwo(t *testing.T) {
	arrayValue := PrimeNumber(2)
	if len(arrayValue) == 0 {
		t.Error("Expect to get empty array!")
	}
}

func TestWithNumberThree(t *testing.T) {
	arrayValue := PrimeNumber(3)
	expected := []int{3}

	if arrayValue[0] != 3 {
		t.Errorf("Got %v\n expected %v", arrayValue, expected)
	}
}

func TestWithNumberFour(t *testing.T) {
	arrayValue := PrimeNumber(4)
	expected := []int{2, 2}

	assert.Equal(t, expected, arrayValue)
}

func TestWithNumberFive(t *testing.T) {
	arrayValue := PrimeNumber(5)
	expected := []int{5}

	assert.Equal(t, expected, arrayValue)
}

func TestWithNumberSix(t *testing.T) {
	arrayValue := PrimeNumber(6)
	expected := []int{3, 2}

	assert.Equal(t, expected, arrayValue)
}

func TestWithNumberSeven(t *testing.T) {
	arrayValue := PrimeNumber(7)
	expected := []int{7}

	assert.Equal(t, expected, arrayValue)
}


func TestWithNumberEight(t *testing.T) {
	arrayValue := PrimeNumber(8)
	expected := []int {2, 2, 2}

	assert.Equal(t, expected, arrayValue)
}

