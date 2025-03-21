package piscine

func NRune(s string, n int) rune {
	if len(s) == 0 || n < 1 || len(s) < n {
		return 0
	}
	srune := []rune(s)
	return srune[n-1]
}
