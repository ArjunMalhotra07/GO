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
	Arrays06()
	Maps07()
	Maps08()
	Structs09()
	Structs10()
	Structs11()
	Structs12()
	If13()
	If14()
	If15()
	Switch16()
	Switch17()
	Loops18()
	Loops19()
	Loops20()
	Loops21()
	Scanner22()
}
