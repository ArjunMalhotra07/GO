//For statements: simple loops, exiting loops,
//looping through collection
package basics

import (
	"fmt"
)

func Loops18() {
	for i := 0; i <= 5; i += 2 {
		fmt.Println(i)
	}
	fmt.Println("*******")
	//i is scoped to the sfor loop only so we wont be getting its value
	fmt.Println("Hlo")
	for i, j := 0, 0; i <= 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println("*******")
	for i := 0; i <= 5; i += 1 {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i *= 2*i + 1
		}
	}
	fmt.Println("*******")
	i := 0
	for ; i <= 5; i += 2 {
		fmt.Println(i)
	}
	fmt.Println(i)
	fmt.Println("*******")
	//i is scoped to the main func so v have access to it
	j := 0
	for j <= 5 {
		fmt.Println(j)
		j += 2
	}
	fmt.Println("*******")
	k := 0
	for {
		fmt.Println(k)
		k += 2
		if k == 10 {
			break
		}
	}
	fmt.Println("*******")
	for a := 1; a <= 10; a += 2 {
		if a%2 == 0 {
			continue
		}
		fmt.Println(a)
	}
}
