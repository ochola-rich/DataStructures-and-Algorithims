package piscine

func ToUpper(s string) string {
	if len(s) < 1 {
		return s
	}
	Upperstr := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v = v - 32
		}
		Upperstr = Upperstr + string(v)
	}
	return Upperstr
}
