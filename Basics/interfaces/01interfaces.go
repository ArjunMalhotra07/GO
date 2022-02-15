package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

//-----------------------------------------
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

//-----------------------------------------
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return c.radius * 4
}

//-----------------------------------------
func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is : %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is : %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		square{length: 7.65},
		circle{radius: 89.2},
		circle{radius: 55.2},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("-----")
	}
}
