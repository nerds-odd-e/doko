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
	b := "3"
	assert.Equal(t, a, b)
}

func inputFizzbuzz(input int64) string {
	return fmt.Sprint(input)
}
