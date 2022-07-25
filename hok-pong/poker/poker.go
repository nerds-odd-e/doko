package poker

import (
	"sort"
	"strings"
)

var rankMap = map[string]int{
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

func Winner(data string) int {
	cards := strings.Split(data, " ")

	rank1, rank2 := ranks(cards[0:5]), ranks(cards[5:])

	sort.Ints(rank1)
	sort.Ints(rank2)

	for i := 4; i >= 0; i-- {
		if rank1[i] < rank2[i] {
			return 2
		}
	}

	return 1

}

func ranks(hand []string) []int {
	ranks := []int{}
	for _, h := range hand {
		ranks = append(ranks, rankMap[string(h[0])])
	}
	return ranks
}
