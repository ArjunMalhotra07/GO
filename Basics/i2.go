package main

import (
	"fmt"
)

func main() {
Loop:
	for a := 1; a <= 10; a += 2 {
		for b := 1; b <= 3; b++ {
			fmt.Println(a * b)
			if a*b > 10 {
				break Loop
			}
		}
	}
}
