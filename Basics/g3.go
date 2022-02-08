/*
MAPS : What r they? Creating. MAnipulation
Structs: ,,          ,,   Naming Convention, Embedding, Tags
*/

package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
	//Field Names + Type associated with that field
}

func main() {
	aDoctor := Doctor{
		number:     3,
		actorName:  "Jon Snow",
		episodes:   []string{},
		companions: []string{"Liz", "Jo", "Sara"},
	}
	fmt.Println(aDoctor)
	//Companions is a stlice of String
	// Structa gathers info that are related to one concept
	//We dont have any constraints on type of data within our struct ie can mix any type of data together
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.companions[2])
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.episodes)
}
