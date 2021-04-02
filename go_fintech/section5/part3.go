package section5

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()

	}
}

func Main3() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		// 10回分Doneを要求させる
		wg.Add(1)
		// producerで値を生成
		go producer(ch, i)
	}
	go consumer(ch, &wg)
	wg.Wait()
	// consumer側からchの受信を待っているためcloseさせる
	close(ch)

}
