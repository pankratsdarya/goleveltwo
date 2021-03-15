package main

import (
	"fmt"
	"sync"
)

const count = 1000

var (
	counter int
	mu      sync.Mutex
	wg      = sync.WaitGroup{}
)

func mainmain() {

	wg.Add(count)

	for i := 0; i < count; i++ {
		go lockUnlock()
	}

	fmt.Println("Wait for it")
	wg.Wait()

	fmt.Println(counter)
}

func lockUnlock() {
	mu.Lock()
	defer mu.Unlock()
	counter++
	wg.Done()
}
