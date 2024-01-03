package basics

import "fmt"

func DoHashing23() {
	performHashingOnArray([]int{1, 10, 50, 1, 10, 100, 200, 50, 99})
}

func performHashingOnArray(nums []int) {
	hashMap := make(map[int]int)
	for _, value := range nums {
		if _, ok := hashMap[value]; ok {
			hashMap[value] += 1
		} else {
			hashMap[value] = 1
		}
	}
	fmt.Println(hashMap)
	fmt.Println(hashMap[15])

}
