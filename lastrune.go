package piscine

func LastRune(s string) rune {
	srune := []rune(s)
	return srune[len(s)-1]
}
