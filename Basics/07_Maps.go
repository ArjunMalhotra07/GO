/*
MAPS : What r they? Creating. MAnipulation
Structs: ,,          ,,   Naming Convention, Embedding, Tags
*/

package basics

import (
	"fmt"
)

func Maps07() {
	//METHOD 1
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
	m := map[[3]int]string{} // ---> This becomes "array" key type, without 3
	//it was slice
	fmt.Println(statePopulations, m)

	//A slice cannot be a key to a map
	//Boolean,strings, pointers, interfaces, structs, arrays, channels

	//METHOD 2
	estatePopulations := make(map[string]int)
	estatePopulations = map[string]int{

		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New york":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(estatePopulations)
	fmt.Println(estatePopulations["Texas"])
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations["Georgia"])
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
}
