package basicsagain

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func TestRoutines() {
	check()
}

var ids = []string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}

func check() {
	t0 := time.Now()
	for i := 0; i < len(ids); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	fmt.Println("Came out of loop")
	wg.Wait()
	fmt.Println("Total Time", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Result from DB", ids[i])
	wg.Done()
}
