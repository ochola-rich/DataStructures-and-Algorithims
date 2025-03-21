package piscine

func Max(a []int) int {
	if len(a) < 1 {
		return 0
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
	return a[len(a)-1]
}
