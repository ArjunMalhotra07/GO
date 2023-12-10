package basics

import (
	"fmt"
)

const myConst int = 100

func Const00() {

	const myConst int16 = 200
	fmt.Printf("%v,%T\n", myConst, myConst)

	const a string = "Hey"
	const b float32 = 3.14
	const c bool = true
	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", c, c)
}

func MainFunction() {
	Const00()
	EnumeratedConstants01()
	Strings02()
	Arrays03()
	Arrays04()
	Arrays05()
}
