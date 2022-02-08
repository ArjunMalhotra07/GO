/* if : Operators, if else, if-else if
Switch: simples cases , casees with multiple test
falling through, type switchees */
package main

import (
	"fmt"
)

func main() {
	var i interface{} = [3]int{}
	switch i.(type) {
	//will return a type
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("Ye b print hoga")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is a [3]int")
	default:
		fmt.Println("i is another type")

	}
}
