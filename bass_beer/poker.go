package main

func AWin(hands string) bool {
	str := hands[len(hands)-2:]
	if getRank(str[:1]) >= 13 {
		return false
	}
	return true
}

func getRank(rankStr string) int {
	return map[string]int{
		"A": 14,
		"K": 13,
	}[rankStr]
}