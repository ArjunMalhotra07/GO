package main

import (
	"fmt"
)

func main() {

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
	for _, v := range statePopulations {
		//set them = range keyword then provide collection that u r going to range over
		fmt.Println(v)

	}
	//if we only want values

	estatePopulations := map[string]int{
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
	for k := range estatePopulations {
		//set them = range keyword then provide collection that u r going to range over
		fmt.Println(k)

	}
	//if we only want keys

}
