package routines

import (
	"fmt"
	"time"
)

func Add() {
	t0 := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go addToArray()
	}
	wg.Wait()
	fmt.Println("Exxecution time", time.Since(t0))
}

func addToArray() {
	var r int
	for i := 0; i < 1000000000; i++ {
		r += 1
	}
	wg.Done()
}
