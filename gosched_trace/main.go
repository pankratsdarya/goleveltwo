package main

import (
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1e8; i++ {
				if i%1e6 == 0 {
					runtime.Gosched()
				}
			}
		}()
	}
	wg.Wait()
}
