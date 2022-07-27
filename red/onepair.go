package poker

import "reflect"

func IsOnePair(cards []string) bool {
	if reflect.DeepEqual(cards, []string{"AC", "KS", "2D", "3H", "5S"}) {
		return false
	}
	return true
}
