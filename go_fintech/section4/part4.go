package section4

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

// interfaceのダッグタイピング

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("get out")
	}
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func Main4() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	DriveCar(mike)
	DriveCar(x)

}
