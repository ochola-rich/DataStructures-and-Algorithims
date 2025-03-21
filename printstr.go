package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	aRune := []rune(s)
	for index := 0; index < len(aRune); index++ {
		z01.PrintRune(aRune[index])
	}
}
