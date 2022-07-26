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

type Hand struct {
	cards []string
}

func get_score(hand []string) int {
	score := 0.0
	for _, card := range hand {
		score = math.Max(score, float64(CARD_SCORE[string(card[0])]))
	}
	return int(score)
}

func three_of_a_kind(hand []string) int {
	count := 0
	for i, card1 := range hand {
		for _, card2 := range hand[i+1:] {
			if string(card1[0]) == string(card2[0]) {
				count += 1
			}
			if count == 2 {
				return CARD_SCORE[string(card1[0])]
			}
		}
	}
	return 0
}

func pair(hand []string) int {
	for i, card1 := range hand {
		for _, card2 := range hand[i+1:] {
			if string(card1[0]) == string(card2[0]) {
				return CARD_SCORE[string(card1[0])]
			}
		}
	}
	return 0
}

func player_one_win(hand1 []string, hand2 []string) bool {
	if three_of_a_kind(hand1) > 0 || three_of_a_kind(hand2) > 0 {
		return three_of_a_kind(hand1) > three_of_a_kind(hand2)
	}
	if pair(hand1) > 0 || pair(hand2) > 0 {
		return pair(hand1) > pair(hand2)
	}
	return get_score(hand1) > get_score(hand2)
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
			if player_one_win(hand1, hand2) {
				score += 1
			}
		}
		return float64(score) / float64(len(games))
	}
	return 0.0
}
