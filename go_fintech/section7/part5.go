package section7

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan<- string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)

	fmt.Println("finish")
	ch <- "result"
}

func Main5() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	go longProcess(ctx, ch)

CTXLOOP:

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
}
