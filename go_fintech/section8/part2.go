package section8

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	Nicknames []string
}

func Main2() {
	b := []byte(`{
		"name": "mike",
		"age": 20,
		"nicknames": ["a", "b", "c"]
	}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)
}
