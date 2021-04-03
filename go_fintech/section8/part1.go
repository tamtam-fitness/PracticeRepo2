package section8

// import (
// 	"fmt"
// 	"net/http"
// 	"net/url"
// )

// func Main1() {
// 	// resp, _ := http.Get("http://example.com")
// 	// defer resp.Body.Close()
// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	// fmt.Println(string(body))
// 	base, _ := url.Parse("http://example.com")
// 	//アクセスしたホスト名以下を問答無用で追加する
// 	reference, _ := url.Parse("/test?a=1b=2")
// 	endpoint := base.ResolveReference(reference).String()
// 	fmt.Println(endpoint)
// 	// postの場合はnilではなくbyteを入れる
// 	req, := http.NewRequest("GET", endpoint, nil)
// 	req,Header.Add("")

// }
