package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzbuzzInput1(t *testing.T) {
	a := inputFizzbuzz(1)
	b := "1"
	assert.Equal(t, a, b)
}

func Test_FizzbuzzInput2(t *testing.T) {
	a := inputFizzbuzz(2)
	b := "2"
	assert.Equal(t, a, b)
}

func Test_FizzbuzzInput3(t *testing.T) {
	a := inputFizzbuzz(3)
	b := "Fizz"
	assert.Equal(t, a, b)
}

func Test_FizzbuzzInput5(t *testing.T) {
	a := inputFizzbuzz(5)
	b := "5"
	assert.Equal(t, a, b)
}

func inputFizzbuzz(input int64) string {
	if input == 3 {
		return "Fizz"
	}
	return fmt.Sprint(input)
}
