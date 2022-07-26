package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player1LooseWithHighCardAgain(t *testing.T) {
	answer := IsPlayer1Win("8C TS KC 9H 4S 7D 2S 5D 4S AC")
	expected := false
	assert.Equal(t, answer, expected, "expected true get %v", answer)
}
