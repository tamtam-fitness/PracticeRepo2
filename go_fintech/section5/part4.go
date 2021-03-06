package section5

import "fmt"

func producer2(first chan<- int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func Main4() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer2(first)
	go multi2(first, second)
	go multi4(second, third)
	for result := range third {
		fmt.Println(result)
	}

}
