package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Olá Mundo") // Goroutine
		waitGroup.Done()
	}()
	go func() {
		write("Programa em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func write(text string) {
	for i := 0; i < 5; {
		fmt.Println(text)
		time.Sleep(time.Second)
		i += 1
	}
}
