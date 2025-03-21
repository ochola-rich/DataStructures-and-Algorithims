package piscine

func IsPrintable(s string) bool {
	if len(s) < 1 {
		return false
	}
	count := 0
	for _, v := range s {
		if v >= ' ' && v < 127 {
			count++
		}
	}
	return count == len(s)
}
