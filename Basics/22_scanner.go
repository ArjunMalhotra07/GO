package basics

import "fmt"

func Scanner22() {
	var a string
	fmt.Println("Enter your 1st name")
	fmt.Scanln(&a)

	var b string
	fmt.Println("Enter your last name")
	fmt.Scanln(&b)

	fmt.Println("Your name is ", a, b)
}
