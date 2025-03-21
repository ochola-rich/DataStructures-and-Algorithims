package piscine

func Join(strs []string, sep string) string {
	resultstr := ""
	for i, substr := range strs {
		resultstr = resultstr + substr
		if i < len(strs)-1 {
			resultstr = resultstr + sep
		}
	}
	return resultstr
}
