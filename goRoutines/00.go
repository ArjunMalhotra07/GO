/* A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program. Or in other words, every concurrently executing activity in Go language is known as a Goroutines. You can consider a Goroutine like a light weighted thread. */
package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {
	// Calling Goroutine
	go display("Welcome")
	display("GeeksforGeeks")
}

/*We added the Sleep() method in our program which makes the main Goroutine sleeps for 1 second in between 1-second the new Goroutine executes, displays “welcome” on the screen, and then terminate after 1-second main Goroutine re-schedule and perform its operation. This process continues until the value of the z<6 after that the main Goroutine terminates. Here, both Goroutine and the normal function work concurrently.
 */
