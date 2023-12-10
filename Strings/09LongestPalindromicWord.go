package strings

import (
	"fmt"
	"strings"
)

func MainFunction09() {

	ansFunc("HeleH Is amma onno is")
}

func ansFunc(str string) {

	var temp byte
	ans := []string{}
	count := 0

	words := strings.Fields(str)

	for i := 0; i < len(words); i++ {
		byte_String := []byte(words[i])

		for i, j := 0, len(byte_String)-1; i < j; i, j = i+1, j-1 {
			temp = byte_String[i]
			byte_String[i] = byte_String[j]
			byte_String[j] = temp
		}
		testWord := string(byte_String)

		if testWord == words[i] {
			ans = append(ans, testWord)
		} else {
			count++
		}
	}

	fmt.Println(len(ans), ans)
	x := fmt.Println
	x(returnAns(ans))
}

func returnAns(ans []string) int {
	max := 0
	for i := 0; i < len(ans); i++ {
		if len(ans[i]) > max {
			max = len(ans[i])
		}
	}
	return max
}
