package piscine

func ShoppingListSort(slice []string) []string {
	for val := range slice {
		for i := 0; i < len(slice)-1-val; i++ {
			if len(slice[i]) > len(slice[i+1]) {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
	return slice
}
