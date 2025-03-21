package piscine

func Atoi(s string) int {
	arune := []rune(s)
	var rint int
	isNegative := false
	if len(s) == 0 {
		return 0
	}
	if arune[0] == '-' || arune[0] == '+' {
		if arune[0] == '-' {
			isNegative = true
		}
		arune = arune[1:]
	}
	for _, c := range arune {
		if c < 48 || c > 57 {
			return 0
		} else {
			nint := int(c) - '0'
			rint = rint*10 + nint
		}
	}
	if isNegative {
		rint = -rint
	}
	return rint
}
