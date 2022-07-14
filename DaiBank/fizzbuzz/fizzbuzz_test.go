package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalNumber(t *testing.T) {
	assert.Equal(t, "1", findFizzBuzz(1))
}
