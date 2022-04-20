package main

import "fmt"

func main() {
	f := fmt.Println
	f("Enter Word to find Duplicate Characters -- ")
	var input string
	fmt.Scanln(&input)
	duplicate1(input)
	testMap := findDuplicates(input)
	f(testMap)
	f(countDuplicates(testMap))
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

func findDuplicates(s string) map[string]int {
	countMap := make(map[string]int)
	count := 0
	testString := []byte(s)
	length := len(s) - 1
	for i := 0; i < len(s)-1; i++ {
		if _, found := countMap[string(testString[i])]; !found {
			for j := i + 1; j < len(s); j++ {
				if testString[i] == testString[j] {
					count += 1
				}
			}
			countMap[string(testString[i])] = count + 1
			count = 0
		}

	}
	if _, found := countMap[string(testString[length])]; !found {
		countMap[string(testString[length])] = 1
	}
	return countMap
}

func countDuplicates(testMap map[string]int) []string {
	ansArray := []string{}

	for letter, count := range testMap {
		if count == 2 {
			ansArray = append(ansArray, letter)
		}
	}

	return ansArray
}
