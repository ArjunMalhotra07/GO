package main

import "fmt"

func main() {
	fmt.Println("Enter Word to find Duplicate Characters -- ")
	var input string
	fmt.Scanln(&input)
	duplicate1(input)
}

func duplicate1(s string) {
	p := fmt.Print
	p1 := fmt.Println
	test := []byte(s)
	ansArray := []byte{}

	for _, ans := range test {
		p(string(ans), " ")

	}
	p1()

	for i := 0; i < len(test)-1; i++ {
		for j := i + 1; j < len(test); j++ {
			if (test[i]) == (test[j]) {
				ansArray = append(ansArray, test[i])
			}
		}

	}
	if len(ansArray) == 0 {
		p1("No Duplicate Characters")
	} else {
		p1("Duplicate Characters -- ", string(ansArray))
	}

}
