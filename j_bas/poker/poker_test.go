package j_bas

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHighCardLoseByHigherRank(t *testing.T) {	
	assert.Equal(t, false, IsPlayerWin("8C TS KC 9H 4S 7D 2S 5D 3S AC"))
}

func TestHighestCardWinByHigherRank(t *testing.T) {	
	assert.Equal(t, true, IsPlayerWin("8C TS AC 9H 4S 7D 2S 5D 3S KC"))
	assert.Equal(t, true, IsPlayerWin("8C TS KC 9H 4S 7D 2S 5D 3S QC"))
	assert.Equal(t, true, IsPlayerWin("8C KS TC 9H 4S 7D 2S 5D 3S QC"))
}

func TestSecondHighestCardLoseByHigherRank(t *testing.T) {	
	assert.Equal(t, false, IsPlayerWin("8C 2C AC 9H 4S TD 2S 5D 3S AS"))
}

func TestSecondHighestCardWinByHigherRank(t *testing.T) {	
	assert.Equal(t, true, IsPlayerWin("8C TS AC 9H 4S 7D 2S 5D 3S AS"))
	//assert.Equal(t, true, IsPlayerWin("8C TS AC 9H 4S 6D 2S 5D 3S AS"))
}

func TestHighestRankOfPlayer1(t *testing.T) {
	assert.Equal(t, 14, HighestRankOfPlayer([]string{ "8C", "TS", "AC", "9H", "4S"}))
	assert.Equal(t, 10, HighestRankOfPlayer([]string{ "8C", "TS", "2C", "9H", "4S"}))
}

func TestSecondRankOfPlayer(t *testing.T) {
	assert.Equal(t, 10, SecondRankOfPlayer([]string{ "8C", "TS", "AC", "9H", "4S"}))
	assert.Equal(t, 6, SecondRankOfPlayer([]string{ "2C", "6S", "AC", "3H", "4S"}))
	assert.Equal(t, 6, SecondRankOfPlayer([]string{ "2A", "6S", "AC", "3H", "4S"}))
}

func TestGetPlayer1Cards(t *testing.T) {
	assert.Equal(t, 
		[]string{ "8C", "TS", "AC", "9H", "4S"}, 
		GetPlayer1Cards("8C TS AC 9H 4S 7D 2S 5D 3S KC"))
	assert.Equal(t, 
		[]string{ "TS", "8C", "AC", "9H", "4S"},
	 	GetPlayer1Cards("TS 8C AC 9H 4S 7D 2S 5D 3S KC"))
}

func TestGetPlayer2Cards(t *testing.T) {
	assert.Equal(t, 
		[]string{ "7D", "2S", "5D", "3S", "KC"}, 
		GetPlayer2Cards("8C TS AC 9H 4S 7D 2S 5D 3S KC"))
}

func TestGetRank(t *testing.T) {
	assert.Equal(t, 2, GetRank("2S"))
	assert.Equal(t, 3, GetRank("3S"))
	assert.Equal(t, 4, GetRank("4S"))
	assert.Equal(t, 10, GetRank("TS"))
	assert.Equal(t, 11, GetRank("JS"))
	assert.Equal(t, 12, GetRank("QS"))
	assert.Equal(t, 13, GetRank("KS"))
	assert.Equal(t, 14, GetRank("AS"))
}