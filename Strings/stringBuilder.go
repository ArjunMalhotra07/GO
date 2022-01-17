package main

import (
	"fmt"
	"strings"
)

func main() {

	var sb strings.Builder
	fmt.Println("This is a string builder", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteString("Hello")
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteString(" World")
	fmt.Println("This is a string builder:", sb.String())
	sb.Grow(100)
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteByte(65)
	fmt.Println("This is a string builder:", sb.String())

}
