package piscine

func ShoppingSummarypositioner(str string) map[string]int {
	position := 0
	res := make(map[string]int)
	for i, substr := range str {
		if substr == ' ' {
			res[str[position:i]] += 1
			position = i + 1
		}
	}
	res[str[position:]] += 1
	return res
}
