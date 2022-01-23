package main

import "fmt"

func main() {
	p := fmt.Println
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	p(restoreString("codeleet", indices))
}

func restoreString(s string, indices []int) string {
	ans := []byte(s)
	for i, index := range indices {
		ans[index] = s[i]
	}
	return string(ans)
}
