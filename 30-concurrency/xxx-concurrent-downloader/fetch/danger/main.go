package main

import (
	"fmt"
	"time"
)

/*
when the default case exists, messages are not guaranteed
*/
func main() {
	ch := make(chan int)

	// if this GR is fast enough
	// the main GR can miss the signals
	go func() {
		var i int

		for {
			i++

			select {
			case ch <- i:
			default:
				// message is lost
				// it doesn't matter whether the chan is buffered or not
			}

			if i == 10000 {
				fmt.Println("gopher dies")
				close(ch)
				return
			}
		}
	}()

	var total int
	for i := range ch {
		// works slower — misses some of the signals
		time.Sleep(time.Nanosecond)
		total += i
	}

	// should be: 50005000
	fmt.Println(total)
}
