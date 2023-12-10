package strings

import (
	"fmt"
	"sort"
)

func MainFunction12() {
	s := []string{"geeks", "for", "geeks", "contribute", "practice"}
	fmt.Println(s)
	min(s, "geeks", "practice")

}

func min(s []string, word1 string, word2 string) {
	p := fmt.Println
	arr2 := []int{}
	arr1 := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == word2 {
			arr2 = append(arr2, i)
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == word1 {
			arr1 = append(arr1, i)
		}
	}
	p(arr2)
	p(arr1)

	minDistance(arr1, arr2)

}

func minDistance(arr1 []int, arr2 []int) {
	minDist := []int{}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {

			testDist := arr1[i] - arr2[j]

			if testDist < 0 {
				testDist *= -1
			}
			minDist = append(minDist, testDist)

		}
	}

	fmt.Println(minDist)
	sort.Ints(minDist)
	fmt.Println(minDist[0])

}
