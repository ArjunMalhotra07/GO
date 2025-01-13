package strings

import (
	"fmt"
	"strconv"
)

func MainFunction01() {
	ans := sum("7A", "3")
	fmt.Println(ans)
}

func sum(str1 string, str2 string) int {

	i, _ := strconv.Atoi(str1)
	j, _ := strconv.Atoi(str2)
	return i + j //gives 3
}
