package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Olá Muno!"
	channel <- "Programando em Go!"

	message := <-channel
	message2 := <-channel
	fmt.Println(message)
	fmt.Println(message2)

}
