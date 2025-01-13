package basics

import (
	"fmt"
)

func Loops20() {
	fmt.Println("Hlo")
	s := []int{100, 2554, 3999}
	for index, value := range s {
		//set them = range keyword then provide collection that u r going to range over
		fmt.Println(index, value)

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
	for index, value := range statePopulations {
		//set them = range keyword then provide collection that u r going to range over
		fmt.Println(index, value)

	}
	d := "Hello Go!!"
	for k, v := range d {
		//set them = range keyword then provide collection that u r going to range over
		fmt.Println(k, string(v))

	}
}
