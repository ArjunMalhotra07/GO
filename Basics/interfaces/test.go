package main

import (
	"fmt"
)

type Student struct {
	name, fatherName, motherName string
	id                           int
}

func main() {
	//f := fmt.Println
	p := fmt.Printf
	a := Student{"Arjun", "Nand Gopal Malhotra", "Suman Mittu", 5788}
	p("\nName: %s", a.name)
	p("\nFathers Name: %s", a.fatherName)
	p("\nMothers Name: %s", a.motherName)
	p("\nID: %d", a.id)
	var name, fName, mName string
	var id int
	fmt.Scanln(&name)
	var b Student
	b.name = name
	fmt.Scanln(&fName)
	b.fatherName = fName
	fmt.Scanln(&mName)
	b.motherName = mName
	fmt.Scanln(&id)
	b.id = id
	p("\nName: %s", b.name)
	p("\nFathers Name: %s", b.fatherName)
	p("\nMothers Name: %s", b.motherName)
	p("\nID: %d", b.id)

}

// func print() {
// 	s := Student{}
// 	p := fmt.Printf
// 	p("\nName: %s", s.name)
// 	p("\nFathers Name: %s", s.fatherName)
// 	p("\nMothers Name: %s", s.motherName)
// }
