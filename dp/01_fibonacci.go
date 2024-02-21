package dp

import "fmt"

func FindFibonacci_01() {
	n := 5
	slice := make([]int, n+1)
	for i := range slice {
		slice[i] = -1
	}
	// fmt.Println(Fibonacci(n, &slice))
	fmt.Println(Fibonacci3(n))

}

func Fibonacci(n int, dpArray *[]int) int {
	if n <= 1 {
		return n
	}
	if (*dpArray)[n] != -1 {
		return (*dpArray)[n]
	}
	(*dpArray)[n] = Fibonacci(n-1, dpArray) + Fibonacci(n-2, dpArray)
	return (*dpArray)[n]
}

func Fibonacci2(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci2(n-1) + Fibonacci2(n-2)
}

func Fibonacci3(n int) int {
	x := 0
	y := 1

	for i := 2; i < n; i++ {
		z := x + y
		x = y
		y = z
	}
	return y
}
