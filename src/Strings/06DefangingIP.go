package strings

import (
	"fmt"
	"strings"
)

func MainFunction06() {

	test := "1.1.1.1"
	fmt.Println(mostString1(test))
}

func mostString1(str string) string {
	test := str
	ans := strings.Replace(test, ".", "[.]", -1)
	return ans
}
