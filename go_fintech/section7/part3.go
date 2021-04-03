package section7

import (
	"fmt"
	"sort"
)

func Main3() {
	p := []struct {
		Name string
		Age  int
	}{
		{"John", 21},
		{"Ken", 18},
		{"Pole", 35},
		{"Nijile", 12},
	}
	fmt.Println(p)
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(p)
}
