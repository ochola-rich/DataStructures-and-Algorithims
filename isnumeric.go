package piscine

func IsNumeric(s string) bool {
	if len(s) < 1 {
		return false
	}
	count := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			count++
		}
	}
	return count == len(s)
}
