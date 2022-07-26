package main

func AWin(hands string) bool {
	str := hands[len(hands)-2:]
	if str[:1] == "A" || str[:1] == "K" {
		return false
	}
	return true
}