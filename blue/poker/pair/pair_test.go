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

// const p2WinPairWithHightCard = "9C 6S KC 5H KS 7D 2S 5D AD AC"

func TestPokerPair(t *testing.T) {
	t.Run("Test Pair P1 Win", func(t *testing.T) {
		assert.Equal(t, true, pair.IsPair(p1WinPair))
	})

	t.Run("Test Pair P2 Win", func(t *testing.T) {
		assert.Equal(t, false, pair.IsPair(p2WinPair))
	})

	t.Run("Test Pair P1 Win2", func(t *testing.T) {
		assert.Equal(t, true, pair.IsPair(p1WinPair2))
	})

	t.Run("TestPairP1WinWithK", func(t *testing.T) {
		assert.Equal(t, true, pair.IsPair(p1WinPairWithK))
	})

	// t.Run("xTestPairP1Lose", func(t *testing.T) {
	// 	assert.Equal(t, false, pair.IsPair(p2WinPairWithHightCard))
	// })
}
