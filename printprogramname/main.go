package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argu := os.Args[0]
	for i := len(argu) - 1; i >= 0; i-- {
		if argu[i] == '/' {
			argu = argu[i+1:]
			break
		}
	}
	for _, v := range argu {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
