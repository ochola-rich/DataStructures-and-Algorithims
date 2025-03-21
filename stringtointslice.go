package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return []int(nil)
	}
	res := []int{}
	for _, char := range str {
		res = append(res, int(char))
	}
	return res
}
