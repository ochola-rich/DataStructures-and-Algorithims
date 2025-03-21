package piscine

func ToLower(s string) string {
	if len(s) < 1 {
		return s
	}
	lowererstr := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			v = v + 32
		}
		lowererstr = lowererstr + string(v)
	}
	return lowererstr
}
