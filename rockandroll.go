package piscine

func RockAndRoll(n int) string {
	str := ""
	switch {
	case n < 0:
		return "error: number is negative\n"
	case n%2 == 0 && n%3 == 0:
		str = "rock and roll\n"
	case n%2 == 0:
		str = "rock\n"
	case n%3 == 0:
		str = "roll\n"
	default:
		str = "error: non divisible\n"
	}
	return str
}
