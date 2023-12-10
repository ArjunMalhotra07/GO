/*
	if : Operators, if else, if-else if

Switch: simples cases , casees with multiple test
falling through, type switchees
*/
package basics

import (
	"fmt"
)

func Switch16() {

	switch 3 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("OThree")
	default:
		fmt.Println("HLOOO")

	}
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("Five, Ten")
	case 2, 4, 6:
		fmt.Println("Two, Four")
	case 3, 7, 9:
		fmt.Println("OThree, Sex")
	default:
		fmt.Println("HLOOO")

	}
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to ten")
		fallthrough
	case i >= 20:
		fmt.Println("greater than or equal to ten")
		//Fallthrough is logic less it does what its told
	default:
		fmt.Println("greater than 20")

	}

}
