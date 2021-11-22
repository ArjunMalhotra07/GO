//ENUMERATED CONSTANTS
package main

import (
	"fmt"
)

const (
	a1 = iota
	b1
	c1
)
const (
	a = 10
	b = iota
	c = iota
)

func main() {
	fmt.Printf("%v,%T\n", a1, a1)
	fmt.Printf("%v,%T\n", b1, b1)
	fmt.Printf("%v,%T\n", c1, c1)
	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", c, c)

}
