package main

func inputPrimefactor(input int64) []int64 {
	if input == 1 {
		return []int64{}
	}

	for factor := int64(2); true; factor++ {
		if input%factor == 0 {
			return append([]int64{factor}, inputPrimefactor(input/factor)...)
		}
	}
	return nil
}
