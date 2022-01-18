package main

import (
	"fmt"
)

func main() {
	ans := check1("125547")
	fmt.Println(ans)
}

func check1(str string) bool {
	b := true
	for _, c := range str {
		if c < '0' || c > '9' {
			b = false
			break
		}
	}
	return b
}
