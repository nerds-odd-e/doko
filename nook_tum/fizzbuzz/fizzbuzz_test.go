package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput1_ShouldSay1(t *testing.T) {
	answer := FizzBuzz(1)
	expected := "1"
	assert.Equal(t, answer, expected, "")
}

func TestInputNumber_ShouldSaySameNumber(t *testing.T) {
	answer := FizzBuzz(4)
	expected := "4"
	assert.Equal(t, answer, expected, "Wanted 4 but got %s", answer)
}

func TestInput1_ShouldSay3(t *testing.T) {
	answer := FizzBuzz(3)
	expected := "Fizz"
	assert.Equal(t, answer, expected, "Fizz")
}

func TestInput6_ShouldSayFizz(t *testing.T) {
	answer := FizzBuzz(6)
	expected := "Fizz"
	assert.Equal(t, answer, expected, "Wanted Fizz but got %s", answer)
}

func TestInput5_ShouldSayBuzz(t *testing.T) {
	answer := FizzBuzz(5)
	expected := "Buzz"
	assert.Equal(t, answer, expected, "Wanted %s but got %s", expected, answer)
}

func TestInput25_ShouldSayBuzz(t *testing.T) {
	answer := FizzBuzz(25)
	expected := "25"
	assert.Equal(t, answer, expected, "Wanted Buzz but got %s", answer)
}

func TestInput15_ShouldSayFizzBuzz(t *testing.T) {
	answer := FizzBuzz(15)
	expected := "FizzBuzz"
	assert.Equal(t, answer, expected, "Wanted FizzBuzz but got %s", answer)
}
