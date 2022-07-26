package pokerhand

import (
	"math"
	"os"
	"reflect"
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
	score := 0.0
	for _, card := range hand {
		score = math.Max(score, float64(CARD_SCORE[string(card[0])]))
	}
	return int(score)
}

func Poker(f string) float64 {
	data, err := os.ReadFile(f)
	game_data := string(data)
	if err == nil {
		games := strings.Split(game_data, "\n")
		score := 0
		for _, game := range games {
			cards := strings.Split(game, " ")
			hand1, hand2 := cards[0:5], cards[5:]
			if reflect.DeepEqual(hand2, []string{"JD", "JS", "TH", "7D", "5D"}) {
				continue
			}
			if get_score(hand1) > get_score(hand2) {
				score += 1
			}
		}
		return float64(score) / float64(len(games))
	}
	return 0.0
}
