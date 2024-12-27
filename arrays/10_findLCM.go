package arrays

import (
	"fmt"
	"math"
)

func DoFindLCM() {
	fmt.Println("Finding LCM")
	fmt.Println(findLCM(5, 20))

}
func findLCM(x, y int) []int {
	gcd := 1
	minValue := int(math.Min(float64(x), float64(y)))
	for i := minValue; i >= minValue/2; i-- {
		if x%i == 0 && y%i == 0 {
			gcd = i
			break
		}
	}
	lcm := gcd * (x / gcd) * (y / gcd)
	return []int{gcd, lcm}
}
