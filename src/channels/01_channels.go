package channels

import "fmt"

func PerformChannels() {
	c := make(chan int)
	go process(c)
	for val := range c {
		fmt.Println(val)
	}
}
func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
