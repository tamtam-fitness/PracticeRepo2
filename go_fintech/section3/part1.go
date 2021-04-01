package section3

import (
	"fmt"
	"os/user"
	"strings"
)

func Main1() {
	fmt.Println(user.Current())
	fmt.Println(string("hello"[0]))

	var s string = "World"
	fmt.Println(strings.Replace(s, "W", "w", 1))
}
