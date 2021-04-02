package section5

import (
	"fmt"
	"time"
)

func goroutine3(ch chan<- string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine4(ch chan<- string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func Main5() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine3(c1)
	go goroutine4(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
