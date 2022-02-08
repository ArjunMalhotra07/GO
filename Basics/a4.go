package main

import (
	"fmt"
)

func main() {

	c := []int{1, 2, 3}
	d := c
	d[2] = 700
	fmt.Println(c)
	fmt.Println(d)
	//This is the slice syntax

	fmt.Printf("Length:%v\n", len(c))
	fmt.Printf("Capacity:%v\n", cap(c))
	//Yaha pe ek slice mei fark daloge sbhi ko pdega

	p := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	q := p[:]
	r := p[:3]
	s := p[3:]
	t := p[3:8]
	p[7] = 57
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println(t)

	/*
		p := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
			q := p[:]
			r := p[:3]
			s := p[3:]
			t := p[3:8]
			p[7] = 57
			fmt.Println(p)
			fmt.Println(q)
			fmt.Println(r)
			fmt.Println(s)
			fmt.Println(t)    iska b same result  ayega coz slicing
			operations do work with arrays as well
	*/

}
