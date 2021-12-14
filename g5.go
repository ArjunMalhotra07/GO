//Structs:  Embedding, Tags
//Instead of inheritance it uses composition
//it supports composition through whats called Embedding

package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

//We can say that Bird has a ima; like characteristics by embedding an animal struct
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

type School struct {
	Name  string
	State string
}

type Student struct {
	School
	studentName string
	id          int
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	c := Student{School: School{Name: "CKCS School", State: "Punjab"},
		studentName: "Arjun M.",
		id:          2021010012673,
	}

	fmt.Println(c)
	fmt.Println(c.studentName)
	//Talking bout modelling behavior, dont use embedding
}
