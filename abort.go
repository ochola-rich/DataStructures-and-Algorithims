package piscine

func Abort(a, b, c, d, e int) int {
	aslice := []int{a, b, c, d, e}
	for i := range aslice {
		for v := 0; v < len(aslice)-1-i; v++ {
			if aslice[v] > aslice[v+1] {
				aslice[v], aslice[v+1] = aslice[v+1], aslice[v]
			}
		}
	}
	return aslice[2]
}
