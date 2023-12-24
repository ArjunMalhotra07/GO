// ! Get Value from a Map
// ! Add/ Edit value to a key in the map
// ! Delete a key from a map
// ! Check if a key exists in a map
// ! Check length of a map
package basicsagain

import "fmt"

func Maps01Basics() {

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
	//! Get value
	fmt.Println(statePopulations["New york"])
	//! Add/Edit value in map
	statePopulations["Georgia"] = 30
	//! Delete a value from Map
	delete(statePopulations, "California")
	fmt.Println(statePopulations)
	//! Check if key exists in map
	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println("Texas Exists")
		fmt.Println(pop, ok)
	} else {
		fmt.Println("Texas ain't here")
	}
	//! Check length of Map variable
	fmt.Println(len(statePopulations))
	//! Assign map to a new map variable
	new_StatePopulation := statePopulations
	delete(new_StatePopulation, "Ohio")
	fmt.Println(statePopulations)
	fmt.Println(new_StatePopulation)
	//Underlyind data is passed by reference
	//Manipulating 1 variable that points to a map is going to have impact on other one
	//! Iterate over maps
	iterateOverMaps(new_StatePopulation)
}

func iterateOverMaps(incomingHashMap map[string]int) {
	for key, value := range incomingHashMap {
		fmt.Println(key, value)
	}
}
