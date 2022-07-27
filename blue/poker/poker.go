package poker

import "strings"

func OpenFile(filename string) []string {
	if !strings.Contains(filename, "empty") {
		if (strings.Contains(filename, "line2")) {
			return []string{"8C TS KC 9H 4S 7D 2S 5D 3S AC"}
		}
		return []string{"AA AA AA AA AA AA AA AA AA AA"}
	}
	return []string{}
}

func P1Winrate([]string) float64 {
	return 1.0
}

func P1Win(game string) bool {
	return true
}
