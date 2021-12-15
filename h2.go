/* if : Operators, if else, if-else if
Switch: simples cases , casees with multiple test
falling through, type switchees */
package main

import (
	"fmt"
)

func main() {

	number := 50
	guess := 105
	if guess < 1 {
		fmt.Println("Guess should be greater than 1 !!")
	} else if guess > 100 {
		fmt.Println("Guess should be less than 100 !!")
	} else {
		if guess < number {
			fmt.Println("too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	//6 diff comparison operarors that we'll use

}

//Concept of short circuiting here, returnTrue test, is a func call but doesnt execute
//coz our compiler didnt check after 1st is true
