package main

import (
	"fmt"
	"strings"
)

func main() {
	x := fmt.Println
	x(reverseLetter("Hey"))

	s := "hi i love u "
	x(reverse_Word(s))

}

func reverseLetter(str string) bool {
	byte_String := []byte(str)
	var temp byte

	for i, j := 0, len(byte_String)-1; i < j; i, j = i+1, j-1 {
		temp = byte_String[i]
		byte_String[i] = byte_String[j]
		byte_String[j] = temp
	}
	test := string(byte_String)
	fmt.Println(test)
	if test == str {
		return true
	} else {
		return false
	}
}

func reverse_Word(s string) bool {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	test := strings.Join(words, " ")

	fmt.Println(test)
	if test == s {
		return true
	} else {
		return false
	}
}
