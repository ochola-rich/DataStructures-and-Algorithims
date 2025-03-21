package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	value := nb
	for i := 2; i <= power; i++ {
		value = value * nb
	}
	return value
}
