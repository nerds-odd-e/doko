package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAWin(t *testing.T) {
	assert.Equal(t, true, AWin("8C TS KC 9H 4S 7D 2S 5D 3S AC")) // in case of A win
}