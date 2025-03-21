package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	res := ""
	count := 0
	result := []string{}
	skip := false
	for _, char := range str {
		if skip {
			skip = false
			continue
		}
		if char != ' ' {
			res += string(char)
			count++
		} else {
			continue
		}
		if count == 5 {
			result = append(result, res)
			res = ""
			count = 0
			skip = true
		}
	}
	if res != "" {
		result = append(result, res)
	}
	chars := ""
	for i, str := range result {
		if i != len(result)-1 {
			chars = chars + str + " "
		} else {
			chars = chars + str
		}
	}
	return chars + "\n"
}
