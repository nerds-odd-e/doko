package pair_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/blue/poker/pair"
)

const p1WinPair = "AC TS KC AH 4S 7D 2S 5D 3S 3C"
const p2WinPair = "8C TS KC 5H 4S 7D AS 5D AD 3C"
const p1WinPair2 = "AC AS KC 5H 4S 7D 2S 5D AD 3C"
const p1WinPairWithK = "9C 6S KC 5H KS 7D 2S 5D AD 3C"

func TestPairP1Win(t *testing.T) {
	assert.Equal(t, true, pair.IsPair(p1WinPair))
}

func TestPairP2Win(t *testing.T) {
	assert.Equal(t, false, pair.IsPair(p2WinPair))
}

func TestPairP1Win2(t *testing.T) {
	assert.Equal(t, true, pair.IsPair(p1WinPair2))
}

func TestPairP1WinWithK(t *testing.T) {
	assert.Equal(t, true, pair.IsPair(p1WinPairWithK))
}

func xTestPairP2win(t *testing.T) {
	assert.Equal(t, false, pair.IsPair(p2WinPair))
}
