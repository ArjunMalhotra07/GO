/* if : Operators, if else, if-else if
Switch: simples cases , casees with multiple test
falling through, type switchees */
package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}
	statePopulations := map[string]int{
		//all keys should be of String type and all value
		//should be of int type
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New york":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	//cannot be single line block in GO

	number := 50
	guess := -50
	if guess < 1 || guess > 100 {
		fmt.Println("Guess should be between 1 and 100 !!")
	}
	if guess >= 1 && guess <= 100 {
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
