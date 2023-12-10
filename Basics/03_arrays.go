package basics

import (
	"fmt"
)

func Arrays03() {
	x := [4]int{78, 100, 54}
	//OR
	z := [...]int{100, 200, 300, 500, 400, 800}
	//means create array which is just large enough
	//to hold that data that i'll pass in the literal syntax
	y := [4]string{"Set", "Wet", "Gel"}
	var students [3]string

	fmt.Printf("Grades are: %v\n", x)
	fmt.Printf("Strings are: %v\n", y)
	fmt.Printf("Grades are: %v\n", z)
	fmt.Printf("Students: %v\n", students)
	students[0] = "Nikit"
	students[1] = "Tommy"
	students[2] = "Singh"
	fmt.Printf("Every student: %v\n", students)
	fmt.Printf("Student 1: %v\n", students[0])
	fmt.Printf("Student 2: %v\n", students[2])
	fmt.Printf("Length: %v\n", len(students)) //Buit in len function

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
	//OR THIS
	var iMatrix [3][3]int
	iMatrix[0] = [3]int{1, 0, 0}
	iMatrix[1] = [3]int{0, 1, 0}
	iMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(iMatrix)

	a := [...]int{1, 2, 3}
	b := a
	b[2] = 700
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	d := &c
	d[2] = 700
	fmt.Println(c)
	fmt.Println(d)

	//Arrays are useful but the fixed length
	//which has to be known at compile time definitely
	//limits its usefulness

}
