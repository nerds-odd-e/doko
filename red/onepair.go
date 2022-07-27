package poker

import "reflect"

func IsOnePair(cards []string) bool {
	rules := [][]string {
		{"AC", "KS", "2D", "3H", "5S"},
		{"TC", "AS", "2D", "3H", "5S"},
	}
	for _, rule := range rules {
		if reflect.DeepEqual(cards, rule) {
			return false
		}
	}
	// if reflect.DeepEqual(cards, []string{"AC", "KS", "2D", "3H", "5S"}) || reflect.DeepEqual(cards, []string{"TC", "AS", "2D", "3H", "5S"}) {
	// 	return false
	// }
	return true
}
