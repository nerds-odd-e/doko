package pokerhand

import (
	"math"
	"os"
	"strings"
)

var CARD_SCORE = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func get_score(hand []string) int {
	score := 0
	for _, card := range hand {
		rank := CARD_SCORE[string(card[0])]
		score = int(math.Max(float64(score), float64(rank)))
	}
	return score
}

func game(hand1 []string, hand2 []string) bool {
	hand1_score := get_score(hand1)
	hand2_score := get_score(hand2)
	return hand1_score > hand2_score

}

func Poker(f string) float64 {
	data, err := os.ReadFile(f)
	cardData := string(data)
	if err == nil {
		cards := strings.Split(cardData, " ")
		hand1, hand2 := cards[0:5], cards[5:]

		if game(hand1, hand2) {
			return 1.0
		} else {
			return 0.0
		}
		// if reflect.DeepEqual(hand1, []string{"KH", "JD", "TH", "7S", "6S"}) {
		// 	return 0.0
		// }
		// return 1.0
	}
	return 0.0
}
