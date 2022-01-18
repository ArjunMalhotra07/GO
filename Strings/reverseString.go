package main

import (
	"fmt"
	"strings"
)

func main() {
	answer := reverse("Hey")
	fmt.Println(answer)
	s := "hello sir big fan"
	ans := reverse_Word(s)
	if ans == s {
		fmt.Println("Yes")
	} else {
		fmt.Println("no")
	}
	words := strings.Fields(s)
	res := strings.Join(words, " ")
	fmt.Println(res)
	for _, ans1 := range s {
		fmt.Printf("%c", ans1)
	}

}

func reverse(str string) string {
	byte_String := []byte(str)
	var temp byte

	for i, j := 0, len(byte_String)-1; i < j; i, j = i+1, j-1 {
		temp = byte_String[i]
		byte_String[i] = byte_String[j]
		byte_String[j] = temp
	}

	return string(byte_String)
}

func reverse_Word(s string) string {

	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
