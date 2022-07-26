package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/pong-chet/fizzbuzz"
)

func Test_FizzBuzz_1_1(t *testing.T) {
	assert.Equal(t, fizzbuzz.FizzBuzz(1), "")
}
