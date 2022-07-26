package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(1)
	want := "1"
	assert.Equal(t, got, want)
}
