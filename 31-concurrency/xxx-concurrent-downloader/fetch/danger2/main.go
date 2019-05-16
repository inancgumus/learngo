package main

import (
	"fmt"
	"time"
)

func main() {
	a, b := make(chan bool), make(chan bool)

	go func() {
		for a != nil || b != nil {
			fmt.Println("loop starts")

			select {
			case <-a:
				fmt.Println("recv: a")
				a = nil
			case <-b:
				b = nil
				fmt.Println("recv: b")
			}

			fmt.Println("loop ends")
		}
		fmt.Println("gopher dies")
	}()

	time.Sleep(time.Second)
	// 	a <- true
	close(a)
	time.Sleep(time.Second)
	//b <- true
	close(b)
	time.Sleep(time.Second * 2)

	// closed chan never blocks
	// nil chan always blocks
	//   if in the loop chans not set to nil, the loop will loop forever
}
