package section7

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func Main6() {
	// content, err := ioutil.ReadFile("./section7/part6.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	r := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	fmt.Println(string(content))
}
