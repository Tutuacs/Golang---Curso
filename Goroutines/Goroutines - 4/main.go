package main

import (
	"fmt"
	"time"
)

func main() {
	chan1, chan2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			chan1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chan2 <- "Canal 2"
		}
	}()

	for {

		// DO THIS
		select {
		case messageChan1 := <-chan1:
			fmt.Println(messageChan1)
		case messageChan2 := <-chan2:
			fmt.Println(messageChan2)
		}

		// INSTEAD OF THIS
		// messageChan1 := <-chan1
		// fmt.Println(messageChan1)

		// messageChan2 := <-chan2
		// fmt.Println(messageChan2)

	}

}
