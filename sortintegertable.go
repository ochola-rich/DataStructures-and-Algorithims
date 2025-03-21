package piscine

func SortIntegerTable(table []int) {
	if len(table) <= 1 {
		return
	}
	for range table {
		for i := 0; i < len(table)-1; i++ {
			if table[i] > table[i+1] {
				table[i], table[i+1] = table[i+1], table[i]
			}
		}
	}
}
