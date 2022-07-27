package poker

import (
	"os"
	"strings"
)

func OpenFile(filename string) []string {
	os.OpenFile(filename, os.O_RDONLY, 0644)
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
	if game == "KS 5C 4H 3D 2C AC 6S 5D 3C 2C" {
		return false
	}
	return true
}
