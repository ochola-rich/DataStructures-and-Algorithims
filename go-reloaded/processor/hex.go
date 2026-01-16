package processor

import (
	"math"
	"strconv"
	"strings"
)

func HextoDec(values []string) {
	s := strings.ToUpper(values[0])
	res := 0

	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		var digit int

		r := runes[i]
		if r >= '0' && r <= '9' {
			digit = int(r - '0')
		} else if r >= 'A' && r <= 'F' {
			digit = int(r-'A') + 10
		}

		power := n - i - 1
		res += digit * int(math.Pow(16, float64(power)))
	}

	values[0] = strconv.Itoa(res)
}
