/*
MAPS : What r they? Creating. MAnipulation
Structs: ,,          ,,   Naming Convention, Embedding, Tags
*/

package main

import "fmt"

func main() {
	aDoctor := struct{ name string }{name: "Jon"} //Anonymous struct
	//1st set of curly brackets is paired to STRUCT keyword and its defining the structure of the struct
	//2nd is the initializer. And its whats goin to provide date into the struct
	//Used when u need to structure data in a way that u dont have in a formal type: short lived
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
	//Unlike maps nd slices these r referring to independent datasets. so it wont change unlike in maps and slices
	//if we wanna point to same underlyind data, we use address of operator
	abcDoctor := &aDoctor
	abcDoctor.name = "Jerry"
	fmt.Println(aDoctor)
	fmt.Println(abcDoctor)

}
