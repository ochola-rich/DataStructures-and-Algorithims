package piscine

func ReverseMenuIndex(menu []string) []string {
	res := make([]string, len(menu))
	index := 0
	for i := len(menu) - 1; i >= 0; i-- {
		res[index] = menu[i]
		index++
	}
	return res
}
