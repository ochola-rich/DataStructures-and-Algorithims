package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 0 {
		return true
	}
	return ascending(f, a) || descending(f, a)
}

func ascending(f func(a, b int) int, a []int) bool {
	for i := range a {
		if i == len(a)-1 {
			break
		}
		if i > 0 && f(a[i], a[i-1]) < 0 {
			return false
		}
	}
	return true
}

func descending(f func(a, b int) int, a []int) bool {
	for i := range a {
		if i == len(a)-1 {
			break
		}
		if i > 0 && f(a[i], a[i-1]) > 0 {
			return false
		}
	}
	return true
}
