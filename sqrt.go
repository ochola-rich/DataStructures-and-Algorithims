package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	for root := 1; root*root <= nb; root++ {
		if root*root == nb {
			return root
		}
	}
	return 0
}
