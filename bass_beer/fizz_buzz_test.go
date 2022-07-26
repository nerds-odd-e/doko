package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnterNumberOne(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
}

func TestEnterNumberTwo(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz(2))
}

func TestEnterNumberThree(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz(3))
}

func TestEnterNumberFour(t *testing.T) {
	assert.Equal(t, "4", fizzbuzz(4))
}

func TestEnterNumberFive(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz(5))
}
