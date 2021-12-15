/* if : Operators, if else, if-else if
Switch: simples cases , casees with multiple test
falling through, type switchees */
package main

import (
	"fmt"
	"math"
)

func main() {

	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	//instead of this use this
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.01 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	//Divide these 2 numbers and if they are within a 10th of a percentage of each other
	//then we'll consider them to be the same
}
