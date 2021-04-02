package section4

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

// Person2 structのNameプロパティを表示させたくない時などに使える
func (p *Person2) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func Main6() {
	mike := &Person2{"MIke", 16}
	fmt.Println(mike)
}
