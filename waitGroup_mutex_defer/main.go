package main

import (
	"os"
	"runtime/trace"
	"sync"
)

const count = 1000

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()

	var (
		counter int
		mu      sync.Mutex
		wg      = sync.WaitGroup{}
	)

	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
			wg.Done()
		}()
	}

	//fmt.Println("Wait for it")
	wg.Wait()

	//fmt.Println(counter)
}
