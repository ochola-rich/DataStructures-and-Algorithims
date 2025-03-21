package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	outputX := [8]rune{'x', ' ', '=', ' ', 52, 50, ',', ' '}
	outputy := [6]rune{'y', ' ', '=', ' ', 50, 49}

	for _, char := range outputX {
		z01.PrintRune(char)
	}
	for _, char := range outputy {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
