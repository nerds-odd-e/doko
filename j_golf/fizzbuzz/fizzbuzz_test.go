package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_1_Return_1(t *testing.T) {
	assert.Equal(t, "1", FizzBuzz(1))
}
