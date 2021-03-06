package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player1LooseWithHighCard_Card5(t *testing.T) {
	answer := IsPlayer1Win("8C TS KC 9H 4S 7D 2S 5D 4S AC")
	expected := false
	assert.Equal(t, expected, answer, "expected player1 loose get %v", answer)
}

func Test_Player1LooseWithHighCard_Card4(t *testing.T) {
	answer := IsPlayer1Win("8C TS KC 9H 4S 7D 2S 5D AS 6C")
	expected := false
	assert.Equal(t, expected, answer, "expected player1 loose get %v", answer)
}

func Test_Player1WinWithHighCard_BothHasA(t *testing.T) {
	answer := IsPlayer1Win("AS TS KC 8C 4S AH 2S 5D 7C 6C")
	expected := true
	assert.Equal(t, expected, answer, "expected player1 win get %v", answer)
}

func Test_Player1WinWithSecondCard(t *testing.T) {
	answer := IsPlayer1Win("AH KH 7C 6H 4S AS TD 8C 7D 4S")
	expected := true
	assert.Equal(t, expected, answer, "expected player1 win get %v", answer)
}

func Test_Player1WinWithHighCard_Card5(t *testing.T) {
	answer := IsPlayer1Win("8C TS KC 9H AS 7D 2S 5D 4S 1C")
	expected := true
	assert.Equal(t, expected, answer, "expected player1 win get %v", answer)
}

func Test_Player1WinWithHighCard_Card4(t *testing.T) {
	answer := IsPlayer1Win("8C TS KC AS 1S 7D 2S 5D 4S 1C")
	expected := true
	assert.Equal(t, expected, answer, "expected player1 win get %v", answer)
}
