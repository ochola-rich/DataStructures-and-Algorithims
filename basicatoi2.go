package piscine

func BasicAtoi2(s string) int {
	arune := []rune(s)
	var rint int
	for _, c := range arune {
		if c < 48 || c > 57 {
			return 0
		} else {
			nint := int(c) - '0'
			rint = rint*10 + nint
		}
	}
	return rint
}
