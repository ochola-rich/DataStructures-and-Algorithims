package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for x := '0'; x <= '9'; x++ {
		for y := '0'; y <= '9'; y++ {
			for i := '0'; i <= '9'; i++ {
				for z := '0'; z <= '9'; z++ {
					if 10*x+y >= 10*i+z {
						continue
					}

					z01.PrintRune(x)
					z01.PrintRune(y)
					z01.PrintRune(' ')
					z01.PrintRune(i)
					z01.PrintRune(z)

					if !(x == '9' && y == '8' && i == '9' && z == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
	z01.PrintRune('\n')
}
