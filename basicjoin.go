package piscine

func BasicJoin(elems []string) string {
	result := ""
	for _, substr := range elems {
		result += substr
	}
	return result
}
