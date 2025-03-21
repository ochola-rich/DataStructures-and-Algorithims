package piscine

func DescendAppendRange(max, min int) []int {
	intslice := []int{}
	if max <= min {
		return intslice
	}
	for i := max; i > min; i-- {
		intslice = append(intslice, i)
	}
	return intslice
}
