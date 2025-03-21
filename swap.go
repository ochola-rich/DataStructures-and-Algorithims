package piscine

func Swap(a *int, b *int) {
	div := *a
	*a = *b
	*b = div
}
