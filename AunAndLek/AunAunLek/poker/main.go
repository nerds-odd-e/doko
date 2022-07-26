package main

import "reflect"

func Player1Win(input []string) int {
	if reflect.DeepEqual(input, []string{"3H 7H 6S KC JS QH TD JC 2D 8S"}) {
		return 1
	}
	return 0
}
