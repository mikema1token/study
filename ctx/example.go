package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	background := context.Background()
	cancel, cancelFunc := context.WithCancel(background)
	go printTask(cancel)
	time.Sleep(time.Second * 10)
	cancelFunc()
}

func printTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("working")
			time.Sleep(time.Second * 1)
		}
	}
}
