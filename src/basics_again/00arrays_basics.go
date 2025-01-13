package basicsagain

import "fmt"

func Arrays00Basics() {
	//! Add to Array
	arr := make([]int, 2)
	// arr := []int{15, 10, 45, 80}
	arr[0] = 10
	arr[1] = 20
	arr = append(arr, 30)
	arr = append(arr, 40)
	fmt.Println(arr)
	//! Add to Slice
	arr1 := []int{}
	arr1 = append(arr1, 6)
	arr1 = append(arr1, 7)
	arr1 = append(arr1, 8)
	fmt.Println(arr1)
	//! Remove element from a Slice
	arr = append(arr[0:2], arr[3:]...)
	fmt.Println(arr)
	//! Pointers to alter ACTUAL value in an array
	usePointers(&arr, 2, 100)
	fmt.Println(arr)
	//! Iterate over array
	iterateOverArray(arr1)
}

func usePointers(array *[]int, index, newValue int) {
	if index >= 0 && index < len(*array) {
		(*array)[index] = newValue
	} else {
		fmt.Println("Invalid index")
	}
}
func iterateOverArray(arr []int) {
	for index, value := range arr {
		fmt.Println(index, value)
	}
	ansArray := [][]int{}
	ansArray = append(ansArray, []int{5, 10})
}
