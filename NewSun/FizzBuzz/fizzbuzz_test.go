package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzbuzzInput1(t *testing.T) {
	a := inputFizzbuzz(1)
	b := ""
	assert.Equal(t, a, b)
}

func inputFizzbuzz(input int64) string {
	return ""
}
