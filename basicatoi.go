package piscine

func BasicAtoi(s string) int {
	arune := []rune(s)
	var rint int

	for _, c := range arune {
		nint := int(c) - '0'
		rint = rint*10 + nint
	}
	return rint
}
