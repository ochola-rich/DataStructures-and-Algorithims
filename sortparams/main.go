package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argu := os.Args[1:]

	for r := range argu {
		for i := 0; i < len(argu)-1-r; i++ {
			if argu[i] > argu[i+1] {
				argu[i], argu[i+1] = argu[i+1], argu[i]
			}
		}
	}

	for _, substr := range argu {
		for _, char := range substr {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
