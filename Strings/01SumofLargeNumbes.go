package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans := sum("7", "3")
	fmt.Println(ans)
}

func sum(str1 string, str2 string) int {

	i, _ := strconv.Atoi(str1)
	j, _ := strconv.Atoi(str2)
	return i + j
}
