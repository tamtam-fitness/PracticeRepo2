package section4

import "fmt"

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func Main3() {
	v := New(3, 4, 5)
	r := v.Area3D()
	fmt.Println(r)
	fmt.Println(v)
}
