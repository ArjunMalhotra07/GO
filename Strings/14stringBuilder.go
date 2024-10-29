package strings

import (
	"fmt"
	"strings"
)

func MainFunction14() {

	var sb strings.Builder
	fmt.Println("This is a string builder", sb.String())

	sb.WriteString("Hello")
	fmt.Println("This is a string builder:", sb.String())

	sb.WriteString(" World")
	fmt.Println("This is a string builder:", sb.String())
	sb.WriteString(" Earth")
	fmt.Println("This is a string builder:", sb.String())
	var myName string = "Arjun"
	var newString string = ""
	for i := 0; i < len(myName); i++ {
		newString += string(myName[i])
	}
	fmt.Println(newString)
	var newSB strings.Builder
	for i := len(myName) - 1; i >= 0; i-- {
		newSB.WriteByte(myName[i])
	}
	myName = newSB.String()
	fmt.Println(myName)

}
