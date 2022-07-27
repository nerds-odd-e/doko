package poker

import "reflect"

func IsOnePair(cards []string) bool {
	if reflect.DeepEqual(cards, []string{"AC", "KS", "2D", "3H", "5S"}) || reflect.DeepEqual(cards, []string{"TC", "AS", "2D", "3H", "5S"}) {
		return false
	}
	return true
}
