package main

import "fmt"

func main() {
	todos := make(chan int, 45)
	result := make(chan int, 45)

	go worker(todos, result)
	go worker(todos, result)
	go worker(todos, result)
	go worker(todos, result)

	for i := 0; i < 45; i++ {
		todos <- i
	}

	close(todos)

	for i := 0; i < 45; i++ {
		result := <-result
		fmt.Println(result)
	}

}

// <-chan Especifica que o canal só recebe dados
// chan<- Especifica que o canal só envia dados
func worker(todos <-chan int, result chan<- int) {
	for todo := range todos {
		result <- fibonacci(todo)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
