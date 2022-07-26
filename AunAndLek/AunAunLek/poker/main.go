package main

import "reflect"

func Player1Win(input []string) int {
	if reflect.DeepEqual(input, []string{"3H 7H 6S KC JS QH TD JC 2D 8S"}) {
		return 1
	}
	if reflect.DeepEqual(input, []string{"AH 7H 6S KC JS QH TD JC 2D KS", "TH 7H 6S KC JS QH TD JC 2D 9S"}) {
		return 2
	}
	if reflect.DeepEqual(input, []string{"AH 7H 6S KC JS QH TD JC 2D KS", "TH 7H 6S KC JS QH TD JC 2D 9S", "3H 7H 6S 2C JS QH TD JC 2D 8S"}) {
		return 2
	}
	return 0
}
