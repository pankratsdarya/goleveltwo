package main

import (
	"fmt"
)

func main() {
	var workers = make(chan struct{}, 1)
	countNum := 0

	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}

		go func(job int) {
			defer func() {
				<-workers
			}()

			fmt.Println(job)
			countNum++
		}(i)
	}

	for i := 0; i < cap(workers); i++ {
		workers <- struct{}{}
	}
	close(workers)

	fmt.Printf("countNum=%d", countNum)
}
