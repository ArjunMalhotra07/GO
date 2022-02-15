package main

import "fmt"

func main() {
	print(5, 10, 15)

	print1("Arjun Malhotra", 2.0, 15)

}

func print(input ...int) {
	fmt.Println(input)
}

func print1(input ...interface{}) {
	fmt.Println(input)
}
