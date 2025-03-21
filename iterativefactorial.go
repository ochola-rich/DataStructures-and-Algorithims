package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb < 1 {
		return 1
	}
	result := 1
	for i := nb; i >= 1; i-- {
		if i > 20 {
			return 0
		}
		result = result * i
	}
	return result
}
