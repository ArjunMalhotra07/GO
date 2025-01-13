package strings

import (
	"fmt"
	"math/big"
)

func MainFunction011() {

	var smallnum, _ = new(big.Int).SetString("2188824200011112223", 10)
	num := smallnum.Uint64()
	fmt.Println(check(num))
}

func check(num uint64) uint64 {
	if num/7 != 0 {
		return 1
	} else {
		return 0
	}
}
