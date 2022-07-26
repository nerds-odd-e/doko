package pokerhand

import "os"

func Poker(f string) float64 {
	_, err := os.ReadFile(f)
	if err == nil {
		return 1.0
	}
	return 0.0
}
