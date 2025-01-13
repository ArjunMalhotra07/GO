package patterns

import "fmt"

func squarePattern_01(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
