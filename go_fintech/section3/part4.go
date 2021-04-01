package section3

import "fmt"

func foo(params ...int) {
	for _, param := range params {
		fmt.Println(param)
	}
}

func Main4() {
	foo(10, 20)
	s := []int{1, 2, 3}
	foo(s...)
}
