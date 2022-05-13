package qing_en_sheen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz1Return1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
}
func TestFizzBuzz2Return2(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz(2))
}

func TestFizzBuzz3ReturnFizz(t *testing.T) {
	assert.Equal(t, "3", fizzbuzz(3))
}

