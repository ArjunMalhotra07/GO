package routines

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var results = []string{}
var m = sync.Mutex{}

func PerformGoRoutines() {
	t0 := time.Now()
	for i := 0; i < len(DbData); i++ {
		wg.Add(1)
		go DbCall(i)
	}
	wg.Wait()
	fmt.Println(results)
	fmt.Println("Exxecution time", time.Since(t0))
}

var DbData = []string{"id1", "id2", "id3", "id4", "id5"}

func DbCall(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("result fromDB", DbData[i])
	m.Lock()
	results = append(results, DbData[i])
	m.Unlock()
	wg.Done()
}
