package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	switch {
	case n < 0:
		return
	case n < 10:
		rnmb := rune(n + '0')
		z01.PrintRune(rnmb)
	}

	dRune := []rune{}
	for n > 0 {
		runedigit := rune(n%10 + '0')
		dRune = append(dRune, runedigit)
		n /= 10
	}

	for range dRune {
		for i := 0; i < len(dRune)-1; i++ {
			if dRune[i] > dRune[i+1] {
				dRune[i], dRune[i+1] = dRune[i+1], dRune[i]
			}
		}
	}
	for i := 0; i < len(dRune); i++ {
		z01.PrintRune(dRune[i])
	}
}
