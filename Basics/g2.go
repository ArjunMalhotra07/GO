/*
MAPS : What r they? Creating. MAnipulation
Structs: ,,          ,,   Naming Convention, Embedding, Tags
*/

package main

import (
	"fmt"
)

func main() {

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{

		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New york":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations["Texas"])
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations["Georgia"])
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations["TEXAS"])
	//Use comma ok syntax to interrogate
	pop, ok := statePopulations["Texas"]
	fmt.Println(pop, ok)
	//Ok was false if the key wasnt found in our map
	//Use it if you r in a situation if you dont know if the key is in the map or not
	fmt.Println(len(statePopulations))
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(statePopulations)
	fmt.Println(sp)
	//Underlyind data is passed by reference
	//Manipulating 1 variable that points to a amp is going to have impact on other one

}
