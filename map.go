package piscine

func Map(f func(int) bool, a []int) []bool {
	res := []bool{}
	for _, val := range a {
		res = append(res, f(val))
	}
	return res
}
