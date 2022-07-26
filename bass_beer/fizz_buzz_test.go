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