package basicsagain

import "fmt"

func MainFunctionOfChannels() {
	// basicChannel()
	MainFunc()
}

func BsicChannel() {
	var c = make(chan int)
	go addToChannel(c)
	for i := range c {
		fmt.Println(i)
	}

}

func addToChannel(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
