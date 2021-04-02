package section5

import (
	"fmt"
	"time"
)

func Main6() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)
OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			// return
			break OuterLoop
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
