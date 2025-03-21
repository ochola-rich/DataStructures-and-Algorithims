package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	argu := os.Args
	buff := make([]byte, 100)
	switch {
	case len(argu) == 1:
		str, _ := os.Stdin.Read(buff)
		for _, char := range string(buff[:str]) {
			z01.PrintRune(char)
		}
	default:
		for i := range argu[1:] {
			filename := argu[1:][i]
			contents, err := os.ReadFile(filename)
			if err != nil {
				printstr("ERROR: " + err.Error())
				os.Exit(1)
			}
			os.Stdout.Write(contents)
		}
	}
}
