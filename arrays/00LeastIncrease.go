package main

import (
	"fmt"
	"math"
)

func main() {
	//f:=fmt.Println
	x := []int{1, 2, 3, 4, 5}
	sum(len(x), x)

}
func sum(length int, arr []int) {
	f := fmt.Println
	f("Input Length:= ", length, " ", "Array:= ", arr)

	sum := 0 /*Is loop mei I am calculating sum of each element in Array using simple For-loop*/
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	f("Sum:= ", sum)                      //Ab sum print kradia
	var sum_float float64 = float64(sum)  //Ab is sum(Jo integer type ka hai) ko float64 type mei convert krdia(decimal: so sum=23 becomes sum=23.00)
	temp := (sum_float / float64(length)) //Ab temp variable mei  is new sum ko divide krdia by length.. take care k humei length b float mei rkhni hogi nhi to error ajega. Kyoki float/int nhi kr skte. ya float/float ye int/int
	if temp == float64(int(temp)) {       //ab is if block mei I am checking k
		f("Minimum Number:= ", temp+1)
	} else {
		f("Minimum Number:= ", math.Ceil(temp))
	}
}
