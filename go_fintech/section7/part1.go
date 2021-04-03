package section7

// time library
import (
	"fmt"
	"time"
)

func Main1() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day())
}
