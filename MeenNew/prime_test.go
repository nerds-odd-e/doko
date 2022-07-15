package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNubmerOne(t *testing.T) {
	Prime := PrimeNumber(1)
	assert.Equal(t, Prime, []int{})
}

func TestNumberTwo(t *testing.T) {
	Prime := PrimeNumber(2)
	assert.Equal(t, Prime, []int{2})
}

func TestNumberThree(t *testing.T) {
	Prime := PrimeNumber(3)
	assert.Equal(t, Prime, []int{3})
}

func TestNumberFour(t *testing.T) {
	Prime := PrimeNumber(4)
	assert.Equal(t, Prime, []int{2, 2})
}

func TestNumberFive(t *testing.T) {
	Prime := PrimeNumber(5)
	assert.Equal(t, Prime, []int{5})
}

func TestNumberSix(t *testing.T) {
	Prime := PrimeNumber(6)
	assert.Equal(t, Prime, []int{2, 3})
}
func TestNumberSeven(t *testing.T) {
	Prime := PrimeNumber(7)
	assert.Equal(t, Prime, []int{7})
}
func TestNumberEight(t *testing.T) {
	Prime := PrimeNumber(8)
	assert.Equal(t, Prime, []int{2, 2, 2})
}

func TestNumberNine(t *testing.T) {
	Prime := PrimeNumber(9)
	assert.Equal(t, Prime, []int{3, 3})
}

func TestNumberTen(t *testing.T) {
	Prime := PrimeNumber(10)
	assert.Equal(t, Prime, []int{2, 5})
}

func TestNumberTwelve(t *testing.T) {
	Prime := PrimeNumber(12)
	assert.Equal(t, Prime, []int{2, 2, 3})
}

// func TestNumber55(t *testing.T) {
// 	Prime := PrimeNumber(55)
// 	assert.Equal(t, Prime, []int{5, 11})
// }
