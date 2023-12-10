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

}
