package piscine

func Capitalize(s string) string {
	str := []rune{}
	upper := true

	for _, v := range s {
		if !(isAlpha(v)) {
			str = append(str, v)
			upper = true
		} else if upper {
			str = append(str, toUpper(v))
			upper = false
		} else {
			str = append(str, toLower(v))
		}
	}
	return string(str)
}

func isAlpha(a rune) bool {
	if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z' || a >= '0' && a <= '9' {
		return true
	}
	return false
}

func toUpper(a rune) rune {
	if a >= 'a' && a <= 'z' {
		a = a - 32
	}
	return a
}

func toLower(a rune) rune {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	return a
}
