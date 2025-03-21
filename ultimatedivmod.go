package piscine

func UltimateDivMod(a *int, b *int) {
	temp := *a % *b
	temp1 := *a / *b
	*a = temp1
	*b = temp
}
