package final_day_poker_hands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestReturn0WinsIfInputIsEmpty(t *testing.T){
	assert.Equal(t,0,pokerhands([]string{}))
}