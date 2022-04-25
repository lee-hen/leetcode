package bulls_and_cows

import (
	"fmt"
)

func GetHint(secret string, guess string) string {
	h := make([]int, 10)

	var bulls, cows int
	n := len(guess)
	for idx := 0; idx < n; idx++ {
		s := secret[idx]
		g := guess[idx]
		if s == g {
			bulls++
		} else {
			if h[s-'0'] < 0 {
				cows++
			}

			if h[g-'0'] > 0 {
				cows++
			}

			h[s-'0']++
			h[g-'0']--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func getHint(secret string, guess string) string {
	lookup := make(map[byte]int)
	var bulls int
	for i := range secret {
		if guess[i] == secret[i] {
			bulls++
		} else {
			lookup[secret[i]]++
		}
	}

	var cows int
	for i := range guess {
		if guess[i] != secret[i] {
			if lookup[guess[i]] != 0 {
				cows++
				lookup[guess[i]]--
			}
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

