package piscine

func IsPrime(nb int) bool {
	// Using the squaroot of  the number method
	// A prime number is a number which has no devisor between 2 and its square root
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
