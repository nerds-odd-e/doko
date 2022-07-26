package main

import "strings"

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

// func getRanksInHand(hand []string) []int {
// 	ranks := []int{}
// 	for _, c := range hand {
// 		ranks = append(ranks, getRank(c))
// 	}
// 	return ranks
// }

func getHighestRankInHand(hand []string) int {
	highestRank := getRank(hand[0])
	for _, c := range hand[1:] {
		if highestRank < getRank(c) {
			highestRank = getRank(c)
		}
	}
	return highestRank
}

func getHands(game string) ([]string, []string) {
	hands := strings.Split(game, " ")
	return hands[:5], hands[5:]
}

func IsP1Win(game string) bool {
	p1Hand, p2Hand := getHands(game)
	return getHighestRankInHand(p1Hand) > getHighestRankInHand(p2Hand)
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
