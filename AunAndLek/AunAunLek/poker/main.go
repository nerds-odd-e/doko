package main

import (
	"sort"
	"strings"
)

func getRank(card string) int {
	ranks := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	return ranks[card[0]]
}

func getRanksInHand(hand []string) []int {
	ranks := []int{}
	for _, c := range hand {
		ranks = append(ranks, getRank(c))
	}
	return ranks
}

func sortRanks(ranks []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))
	return ranks
}

func getHands(game string) ([]string, []string) {
	hands := strings.Split(game, " ")
	return hands[:5], hands[5:]
}

func IsP1Win(game string) bool {
	p1Hand, p2Hand := getHands(game)
	p1Ranks := sortRanks(getRanksInHand(p1Hand))
	p2Ranks := sortRanks(getRanksInHand(p2Hand))
	return p1Ranks[0] > p2Ranks[0]
}

func Player1WinCount(games []string) int {
	p1WinCount := 0
	for i := 0; i < len(games); i++ {
		if IsP1Win(games[i]) {
			p1WinCount = p1WinCount + 1
		}
	}
	return p1WinCount
}
