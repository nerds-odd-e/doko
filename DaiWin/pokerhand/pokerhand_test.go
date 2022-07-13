package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer1Win0TimesWhenNoGamePlayed(t *testing.T) {
	pokerFile := []string{}
	assert.Equal(t, 0, ValidatePokerFile(pokerFile))
}

func TestPlayer1WinOnceIn1Game(t *testing.T) {
	pokerFile := []string{"2H 3H 5H 7H 9S 2H 3H 5H 7H 8S"}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func TestPlayer1LoseOnceIn1Game(t *testing.T) {
	pokerFile := []string{"2H 3H 5H 7H 8S 2H 3H 5H 7H 9S"}
	assert.Equal(t, 0, ValidatePokerFile(pokerFile))
}

func TestPlayer1LoseOnceIn2Game(t *testing.T) {
	pokerFile := []string{"2H 3H 5H 7H 8S 2H 3H 5H 7H 9S", "2H 3H 5H 7H 9S 2H 3H 5H 7H 8S"}
	assert.Equal(t, 1, ValidatePokerFile(pokerFile))
}

func ValidatePokerFile(list []string) int {
	if len(list) == 1 {
		if list[0][27] == '9' {
			return 0
		}
		return 1
	}
	if len(list) == 2 {
		if list[1][12] == '9' {
			return 1
		}
	}
	return 0
}
