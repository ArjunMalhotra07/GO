package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}() //Tese parenthesis are executing this func
	//This is an anonymous func: one that doesn't have a name
	//Nothing else can call it, it is defined once and called once
	//Recover func will return nil if the func isnt panicking, if not then it will return the error
	panic("Something bad happened")
	fmt.Println("End")
}
