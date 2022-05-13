package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturn0WinsIfInputIsEmpty(t *testing.T) {
	assert.Equal(t, 0, pokerhands([]string{}))
}


func TestReturn1IfP1HasHighestCard(t *testing.T){
	assert.Equal(t,0,pokerhands([]string{"8C TS KC 9H AS 4S 7D 2S 5D 3S"}))
}
