package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	count := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' {
			count++
		}
	}
	return count == len(s)
}
