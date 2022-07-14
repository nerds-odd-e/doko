package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const unSortCard = "2H 3H TS 4H 5H"
const unSortCardWithTenAndKing = "2H 3H TS KH 5H"

func TestHighPointCardInHand(t *testing.T) {
	assert.Equal(t, 10, findHighCardPointInHand(unSortCard))
}

func TestHighPointCardInHandWithTenAndKing(t *testing.T) {
	assert.Equal(t, 13, findHighCardPointInHand(unSortCardWithTenAndKing))
}

func findHighCardPointInHand(hand string) int {
	mapPointWithHonorCard := map[string]string{
		"A": "14",
		"K": "13",
		"Q": "12",
		"J": "11",
		"T": "10",
		"9": "9",
		"8": "8",
		"7": "7",
		"6": "6",
		"5": "5",
		"4": "4",
		"3": "3",
		"2": "2",
	}

	arrCardInHand := strings.Split(hand, " ")
	highCard, _ := strconv.Atoi(mapPointWithHonorCard[string(strings.Split(hand, " ")[0][0])])
	for _, cardInHand := range arrCardInHand {
		pointCard, _ := strconv.Atoi(mapPointWithHonorCard[string(cardInHand[0])])
		if highCard < pointCard {
			highCard = pointCard
		}

	}
	return highCard
}
