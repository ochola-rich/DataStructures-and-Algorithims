package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argu := os.Args[1:]

	for i := len(argu) - 1; i >= 0; i-- {
		for _, char := range argu[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
