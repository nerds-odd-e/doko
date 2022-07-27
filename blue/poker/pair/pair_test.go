package pair_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/blue/poker/pair"
)

const p1WinPair = ""

func TestPairP1Win(t *testing.T) {
	assert.Equal(t, true, pair.IsPair(p1WinPair))
}
