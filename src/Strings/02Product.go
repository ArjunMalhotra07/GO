package strings

import (
	"fmt"
	"strconv"
)

func MainFunction02() {
	ans := Product("7", "3")
	fmt.Println(ans)
}

func Product(str1 string, str2 string) int {

	i, _ := strconv.Atoi(str1)
	j, _ := strconv.Atoi(str2)
	return i * j
}
