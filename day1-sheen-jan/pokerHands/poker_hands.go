package pokerHands

import (
	"strconv"
	"strings"
)

var RankMap = map[string]int{
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func pokerHands(hands []string) int {
	player1Points := 0
	for _, hand := range hands {
		cards := strings.Split(hand, " ")

		p1HighestRank, p1HighestSuite := getHighestCard(cards[:5])
		p2HighestRank, p2HighestSuite := getHighestCard(cards[5:])

		if p1HighestRank > p2HighestRank {
			player1Points += 1
		}
		if p1HighestRank == p2HighestRank && p1HighestSuite > p2HighestSuite {
			player1Points += 1
		}
	}
	return player1Points
}

func getHighestCard(hand []string) (int, string) {
	currentHighestRank := 0
	highestSuite := ""
	for _, card := range hand {
		cardRankStr := string(card[0])
		cardRank, err := strconv.Atoi(cardRankStr)

		if err != nil {
			cardRank = RankMap[cardRankStr]
		}
		if cardRank > currentHighestRank {
			currentHighestRank = cardRank
			highestSuite = string(card[1])
		}
	}
	return currentHighestRank, highestSuite
}
