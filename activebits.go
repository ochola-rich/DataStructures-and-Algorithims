package piscine

func ActiveBits(n int) int {
	intval := []int{}
	count := 0
	for n != 0 {
		val := n % 2
		intval = append(intval, val)
		n /= 2
	}
	for _, val := range intval {
		if val == 1 {
			count++
		}
	}
	return count
}
