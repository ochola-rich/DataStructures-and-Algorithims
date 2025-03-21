package piscine

func Rot14(s string) string {
	resultstr := ""

	for _, char := range s {
		if char >= 'a' && char <= 'l' || char >= 'A' && char <= 'L' {
			resultstr = resultstr + string(char+14)
		} else if char >= 'm' && char <= 'z' || char >= 'M' && char <= 'Z' {
			resultstr = resultstr + string(char-12)
		} else {
			resultstr += string(char)
		}
	}
	return resultstr
}
