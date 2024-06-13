package basicsagain

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RoutinesMainFunction() {
	// CheckGoRoutine()
	WriteToSharedMemorySpace()
}

var ids = []string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}
var ans = []string{}
var mutex = sync.RWMutex{}

func CheckGoRoutine() {
	t0 := time.Now()
	for i := 0; i < len(ids); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	fmt.Println("Came out of loop")
	wg.Wait()
	fmt.Println("Total Time", time.Since(t0))
	fmt.Println("Received Values:", ans)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Result from DB", ids[i])
	writeCall(ids[i])
	readLock()
	wg.Done()
}

func writeCall(result string) {
	mutex.Lock()
	ans = append(ans, result)
	mutex.Unlock()
}
func readLock() {
	mutex.RLock()
	fmt.Printf("Current Results : %v\n\n", ans)
	mutex.RUnlock()
}
