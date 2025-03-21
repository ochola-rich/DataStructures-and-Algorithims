package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, substr := range a {
		for _, char := range substr {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
