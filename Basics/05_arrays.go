package basics

import (
	"fmt"
)

func Arrays05() {

	c := make([]int, 3, 100)
	//make is built-in function to create slices that can take 2 or 3 arguments
	//3rd argument is CAPACITY
	fmt.Println(c)
	fmt.Printf("Length:%v\n", len(c))
	fmt.Printf("Capacity:%v\n", cap(c))

	a := make([]int, 3)
	fmt.Println(a)
	fmt.Printf("Length:%v\n", len(a))
	fmt.Printf("Capacity:%v\n", cap(a))

	b := []int{}
	fmt.Println(b)
	fmt.Printf("Length:%v\n", len(b))
	fmt.Printf("Capacity:%v\n", cap(b))

	b = append(b, 100)
	fmt.Println(b)
	fmt.Printf("Length:%v\n", len(b))
	fmt.Printf("Capacity:%v\n", cap(b))

	b = append(b, []int{200, 300, 400, 50, 60}...)
	fmt.Println(b)
	fmt.Printf("Length:%v\n", len(b))
	fmt.Printf("Capacity:%v\n", cap(b))

}
