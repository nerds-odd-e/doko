package main

import (
	"sort"
	"strings"
)

func ComparePokerHand(game string) string {
	gameArr := strings.Split(game, " ")
	p1 := gameArr[0:5]
	p2 := gameArr[5:]

	if getHighCard(p1) == getHighCard(p2) {
		if getSecondHighCard(p1) > getSecondHighCard(p2) {
			return "win"
		}
		return "lose"
	}

	if getHighCard(p1) > getHighCard(p2) {
		return "win"
	}

	return "lose"
}

func getHighCard(hand []string) int {
	return getNthHighCard(hand, 1)
}

func getSecondHighCard(hand []string) int {
	return getNthHighCard(hand, 2)
}

func getNthHighCard(hand []string, nth int) int {
	scoreMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	cardScores := []int{}
	for _, card := range hand {
		cardScores = append(cardScores, scoreMap[string(card[0])])
	}

	sort.Ints(cardScores)
	return cardScores[len(cardScores)-nth]
}
