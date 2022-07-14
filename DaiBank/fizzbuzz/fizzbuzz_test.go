package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizz(t *testing.T) {
	assert.Equal(t, "BUZZ", findFizzBuzz(1))
}
