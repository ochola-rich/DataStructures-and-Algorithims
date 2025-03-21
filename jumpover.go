package piscine

func JumpOver(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 3 {
		return "\n"
	}
	res := ""
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	return res + "\n"
}
