package strings

import (
	"fmt"
)

func MainFunction04() {
	test := "Arjun"
	ans := reverse(test, len(test)-1)
	fmt.Println(ans)

}

func reverse(s string, length int) string {
	ans := length

	byte_String := []byte(s)
	ans_String := []byte{}

	for ans >= 0 {
		ans_String = append(ans_String, byte_String[ans])
		return string(ans_String) + reverse(s, ans-1)
	}

	return ""
}
