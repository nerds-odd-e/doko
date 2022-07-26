package main

func AWin(hands string) bool {
	str := hands[len(hands)-2:]
	if getRank(str[:1]) >= 13 {
		return false
	}
	return true
}

func getRank(rankStr string) int {
	if rankStr == "A" {
		return 14
	}
	if rankStr == "K" {
		return 13
	}
	return 0
}