package strings

import "fmt"

func DoReverseArray() {
	testArray := []int{5, 10, 15, 20}
	ReverseArray(&testArray, 0)
	fmt.Println(testArray)
}

func ReverseArray(a *[]int, index int) {
	if index >= len(*a)/2 {
		return
	}
	temp := (*a)[index]
	(*a)[index] = (*a)[len(*a)-index-1]
	(*a)[len(*a)-index-1] = temp
	ReverseArray(a, index+1)
}
