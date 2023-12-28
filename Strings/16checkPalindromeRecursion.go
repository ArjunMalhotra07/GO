package strings

func CheckRecursionUsingRecursion(n string, index int) bool {
	if index >= len(n)/2 {
		return true
	}
	if n[index] != n[len(n)-index-1] {
		return false
	}
	return CheckRecursionUsingRecursion(n, index+1)
}

