package pair_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/blue/poker/pair"
)

const p1WinPair = "AC TS KC AH 4S 7D 2S 5D 3S 3C"
const p2WinPair = "8C TS KC 5H 4S 7D AS 5D AD 3C"

func TestPairP1Win(t *testing.T) {
	assert.Equal(t, true, pair.IsPair(p1WinPair))
}

func TestPairP2Win(t *testing.T) {
	assert.Equal(t, false, pair.IsPair(p2WinPair))
}
