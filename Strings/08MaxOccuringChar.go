//-->How to find the maximum occurring character in given String?
package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	p(testCheck("hello"))

}
func testCheck(str string) string {
	p := fmt.Println
	map1 := make(map[string]int)
	for i := 0; i < len(str); i++ {
		tempS := string(str[i])
		tempI := strings.Count(str, tempS)

		map1[tempS] = tempI
	}
	p(map1)
	max := 0
	ans := ""

	count := 0
	for key, value := range map1 {
		if value > max {
			max = value
			ans = key
		} else {
			count++
		}
	}
	return ans

}
