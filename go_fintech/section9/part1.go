package section9

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	isAquire := s.TryAcquire(1)
	if !isAquire {
		fmt.Println("Could not get lock")
		return
	}

	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	defer s.Release(1)
	fmt.Println("Wait....")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func Main1() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)

	time.Sleep(5 * time.Second)
	go longProcess(ctx)
}
