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

func (h Hand) high_card() int {
	score := 0.0
	for _, card := range h.cards {
		score = math.Max(score, float64(CARD_SCORE[string(card[0])]))
	}
	return int(score)
}

func (h Hand) three_of_a_kind() int {
	count := 0
	for i, card1 := range h.cards {
		for _, card2 := range h.cards[i+1:] {
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

func (h Hand) pair() int {
	for i, card1 := range h.cards {
		for _, card2 := range h.cards[i+1:] {
			if string(card1[0]) == string(card2[0]) {
				return CARD_SCORE[string(card1[0])]
			}
		}
	}
	return 0
}

func player_one_win(hand1 Hand, hand2 Hand) bool {
	if hand1.three_of_a_kind() > 0 || hand2.three_of_a_kind() > 0 {
		return hand1.three_of_a_kind() > hand2.three_of_a_kind()
	}
	if hand1.pair() > 0 || hand2.pair() > 0 {
		return hand1.pair() > hand2.pair()
	}
	return hand1.high_card() > hand2.high_card()
}

func Poker(f string) float64 {
	data, err := os.ReadFile(f)
	game_data := string(data)
	if err == nil {
		games := strings.Split(game_data, "\n")
		score := 0
		for _, game := range games {
			cards := strings.Split(game, " ")
			hand1, hand2 := Hand{cards[0:5]}, Hand{cards[5:]}
			if player_one_win(hand1, hand2) {
				score += 1
			}
		}
		return float64(score) / float64(len(games))
	}
	return 0.0
}
