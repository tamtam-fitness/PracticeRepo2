package section5

import "fmt"

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Main2() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine2(s, c)
	x := <-c
	fmt.Println(x)
}
