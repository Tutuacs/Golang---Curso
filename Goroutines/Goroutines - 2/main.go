package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go write("Olá Mundo!", channel)

	// DO THIS
	for message := range channel {
		fmt.Println(message)
	}

	// NOT THIS
	// for {
	// 	message, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	fmt.Println("Fim do Programa!")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
