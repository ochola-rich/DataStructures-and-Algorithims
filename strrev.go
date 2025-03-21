package piscine

func StrRev(s string) string {
	var strg []byte
	for i := (len(s) - 1); i >= 0; i-- {
		strg = append(strg, s[i])
	}
	return string(strg)
}
