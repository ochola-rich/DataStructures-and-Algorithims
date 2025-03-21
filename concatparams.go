package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, substr := range args {
		result = result + substr
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
