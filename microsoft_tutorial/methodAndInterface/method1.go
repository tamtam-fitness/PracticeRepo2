package methodAndInterface

type Triangle struct {
	size int
}

func (t *Triangle) doubleSize() {
	t.size *= 2
}

func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}

// import (
// 	"fmt"

// 	"example.com/me/hello/methodAndInterface"
// )

// func main() {
// 	t := methodAndInterface.Triangle{}
// 	t.SetSize(3)
// プライベートなのでエラーになる
// 	fmt.Println(t.size)

// 	fmt.Println(t.Perimeter())
// }
