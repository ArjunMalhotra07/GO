package arrays

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == b {
		return a
	}
	if a > b {
		return GCD(a-b, b)
	}
	return GCD(a, b-a)
}
