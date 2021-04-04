package section9

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func Main3() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
