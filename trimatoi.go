package piscine

func TrimAtoi(s string) int {
	arune := []rune(s)
	number := 0
	index := []int{}
	for i, v := range arune {
		if v == '-' {
			index = append(index, i)
		}
		if v >= '0' && v <= '9' {
			number = (number * 10) + int(v-'0')
			index = append(index, i)
		}
	}
	if number <= 0 {
		return 0
	} else if arune[index[0]] == '-' {
		number = -number
	}
	return number
}
