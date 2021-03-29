package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// コンカレンシーを実装する改良バージョン。 現時点では、かかる時間は数秒 (15 秒以内)のはずです。 バッファーありのチャネルを使用する必要があります。

func fib(number float64, ch chan<- float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- x
}

func Task1() {
	start := time.Now()

	ch := make(chan float64, 16)
	for i := 1; i < 15; i++ {
		go fib(float64(i), ch)
	}

	for i := 0; i < 14; i++ {
		fmt.Printf("Fib(%v): %v\n", i, <-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
