package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("app starting")

	wasSignal := false
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for !wasSignal {
			fmt.Println("awaiting signal")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	sig := <-sigs
	wasSignal = true
	fmt.Println(sig)
	fmt.Println("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		for i := 9; i >= 0; i-- {
			fmt.Printf("0.%d \n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	<-ctx.Done()
	fmt.Println("app closed")
}
