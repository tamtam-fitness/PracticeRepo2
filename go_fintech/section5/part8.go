package section5

import "fmt"

func goroutine5(s []string, c chan<- string) {
	defer close(c)

	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}

}

func Main8() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine5(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
