package piscine

func Unmatch(a []int) int {
	for i := range a {
		count := 0
		for _, v := range a {
			if a[i] == v {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}
