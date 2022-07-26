package main

func AWin(hands string) bool {
	str := hands[len(hands)-2:]
	if str == "AC" {
		return false
	}
	return true
}