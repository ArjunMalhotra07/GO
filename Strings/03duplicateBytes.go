package main

import "fmt"

func main() {
	duplicate("sunaina")
}

func duplicate(s string) {
	p := fmt.Print
	p1 := fmt.Println
	test := []byte(s)
	ansArray := []byte{}

	for _, ans := range test {
		p(string(ans))
		p1()
	}
	p1()

	for i := 0; i < len(test)-1; i++ {
		//count := 0
		for j := i + 1; j < len(test); j++ {
			if string(test[i]) == string(test[j]) {
				ansArray = append(ansArray, test[i])
			}
		}

	}
	p1(string(ansArray))
}
