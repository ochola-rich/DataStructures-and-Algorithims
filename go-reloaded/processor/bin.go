package processor

import (
	"math"
	"strconv"
)

func BintoDec(values []string) {
	res := 0

	runes := []rune(values[0])
	n := len(runes)

	for i := 0; i < n; i++ {
		var digit int

		r := runes[i]
		if r >= '0' && r <= '1' {
			digit = int(r - '0')
		}
		power := n - i - 1
		res += digit * int(math.Pow(2, float64(power)))
	}

	values[0] = strconv.Itoa(res)
}
