//Stack operations
package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5}
	//POP OFF from the stack
	b := a[:len(a)-1]
	fmt.Println(b)
	c := append(a[:2], a[3:]...)
	fmt.Println(c)
	fmt.Println(a)

}
