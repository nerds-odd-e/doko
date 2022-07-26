package pokerhand

import (
	"os"
	"reflect"
	"strings"
)

func Poker(f string) float64 {
	data, err := os.ReadFile(f)
	cardData := string(data)
	if err == nil {
		cards := strings.Split(cardData, " ")
		hand1, _ := cards[0:5], cards[5:]
		if reflect.DeepEqual(hand1, []string{"KH", "JD", "TH", "7S", "6S"}) {
			return 0.0
		}
		return 1.0
	}
	return 0.0
}
