package basicsagain

import (
	"fmt"
	"sync"
)

var counter int
var m = sync.Mutex{}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		m.Lock()
		counter++
		m.Unlock()
	}
}

func WriteToSharedMemorySpace() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
