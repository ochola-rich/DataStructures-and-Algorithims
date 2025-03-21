package piscine

func SplitWhiteSpaces(s string) []string {
	resulstr := []string{}
	substr := ""
	for _, char := range s {
		if !(char == ' ' || char == '\n' || char == '\t') {
			substr = substr + string(char)
		} else {
			if substr != "" {
				resulstr = append(resulstr, substr)
				substr = ""
			}
		}
	}
	if substr != "" {
		resulstr = append(resulstr, substr)
	}
	return resulstr
}
