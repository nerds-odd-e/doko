package meennew

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test35(t *testing.T) {
	test35 := FizzBuzz(15)
	assert.Equal(t, test35, "FIZZBUZZ")
}

func Test3(t *testing.T) {
	test3 := FizzBuzz(3)
	assert.Equal(t, test3, "FIZZ")
}
func Test5(t *testing.T) {
	test5 := FizzBuzz(5)
	assert.Equal(t, test5, "BUZZ")
}

func TestNoMatchNumberTwo(t *testing.T) {
	test2 := FizzBuzz(2)
	assert.Equal(t, test2, "2")
}
func TestNoMatchNumberOne(t *testing.T) {
	test1 := FizzBuzz(1)
	assert.Equal(t, test1, "1")
}

func TestFizzBuzzCaseTwo(t *testing.T) {
	testFizzBuzz := FizzBuzz(30)
	assert.Equal(t, testFizzBuzz, "FIZZBUZZ")
}
