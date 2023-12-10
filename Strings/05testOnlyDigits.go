package strings

import (
	"fmt"
)

func MainFunction05() {
	ans := check1("125547")
	fmt.Println(ans)

	s := "fkf2929"
	ans1 := hlo(s)
	fmt.Println(ans1)
	fmt.Println([]byte(s))
}

func check1(str string) bool {
	b := true
	for _, c := range str {
		if c < '0' || c > '9' {
			b = false
			break
		}
	}
	return b
}

// OR
func hlo(str string) bool {
	test := []byte(str)

	for i := 0; i < len(str); i++ {
		if (test[i]) <= 48 || (test[i]) >= 57 {
			return false
		}
	}
	return true
}
