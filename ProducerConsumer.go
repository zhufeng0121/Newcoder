package main

import (
	"fmt"
	"time"
)


// chan <- 只发送通道
func Producer(begin, end int, queue chan <- int) {
	for i := begin; i < end; i++ {
		fmt.Println("producer", i)
		queue <- i
	}
}

func Consumer(queue <- chan int ) {
	for val := range queue {
		fmt.Println("consumer",val)
	}
}

func main() {
	queue := make(chan int)
	defer close(queue)

	for i := 0; i < 3; i++ {
		go Producer(i * 5, (i + 1) * 5,queue)
	}
	go Consumer(queue)

	time.Sleep(time.Second)
}

