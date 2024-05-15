package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write("Olá Mundo!"), write("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}

func multiplex(inChan, inChan2 <-chan string) <-chan string {
	outChan := make(chan string)

	go func() {
		for {
			select {
			case msg := <-inChan:
				outChan <- msg
			case msg := <-inChan2:
				outChan <- msg
			}
		}
	}()

	return outChan
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
