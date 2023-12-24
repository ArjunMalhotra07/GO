package patterns

import "fmt"

func leftRightTriangle_02(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func leftRightTriangle_03(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

func leftRightTriangle_04(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(i + 1)
		}
		fmt.Println()
	}
}

func leftRightTriangle_05(n int) {
	for i := 0; i <= n; i++ {
		for j := n - i + 1; j > 0; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
