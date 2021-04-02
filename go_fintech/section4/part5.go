package section4

import "fmt"

// interfaceをstringやintにする -> type-assertion
// int をfloatにする -> type convertion or cast

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("Idon't know %T\n", v)
	}
}

func Main5() {
	do(10)
	do("Mike")
	do(true)
}
