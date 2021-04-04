package section8

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json: "-"`
	Age       int      `json: "age, omitempty"`
	Nicknames []string `json: "nicknames"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	// a := struct{Name string}{Name: "test"}
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
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

	v, _ := json.Marshal(p)
	fmt.Println(string(v))

}
