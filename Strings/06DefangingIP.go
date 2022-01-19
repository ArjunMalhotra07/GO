package main

import (
	"fmt"
	"strings"
)

func main() {

	test := "1.1.1.1"
	fmt.Println(mostString1(test))
}

func mostString1(str string) string {
	test := str
	ans := strings.Replace(test, ".", "[.]", -1)
	return ans
}
