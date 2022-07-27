package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_1_Return_1(t *testing.T) {
	input := 1
	expect := "1"

	actual := FizzBuzz(input)

	assert.Equal(t, expect, actual)
}
