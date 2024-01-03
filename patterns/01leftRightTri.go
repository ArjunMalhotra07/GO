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
		//! [Total - rowNumber + 1] times
		for j := 0; j < n-i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func leftRightTriangle_06(n int) {
	for i := 0; i <= n; i++ {
		//! [Total - rowNumber + 1] times
		for j := 0; j < n-i+1; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

func uprightPyramid_07(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func downwardPyramid_08(n int) {
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
func pyramid_09(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
