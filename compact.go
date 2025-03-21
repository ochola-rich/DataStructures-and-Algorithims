package piscine

func Compact(ptr *[]string) int {
	res := []string{}
	for _, char := range *ptr {
		if char == "" {
			continue
		}
		res = append(res, char)
	}
	*ptr = res
	return len(res)
}
