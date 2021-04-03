package section7

import (
	"fmt"
	"regexp"
)

func Main2() {
	//  aとeがあり、間に[a-z]の文字列が1文字以上
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	// 正規表現
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// s := "/view/test"
	// ^は先頭という意味 $は何かしらアルフベットを含んだもの
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/edit/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}
