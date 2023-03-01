package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	quit := make(chan int)
	count := 10
	var wg sync.WaitGroup

	wg.Add(1)
	go ping(count, c1, c2, quit)
	go pong(count, c1, c2, quit)

	c1 <- 0
	res := <-quit
	if res > 0 {
		fmt.Println("res:", res)
		wg.Done()

	}
	wg.Wait()
	close(c1)
	close(c2)
	close(quit)

}

func ping(count int, c1, c2, quit chan int) {
	for {
		time.Sleep(2 * time.Second)
		x := <-c1
		fmt.Println("ping:", x)
		if x < count {
			c2 <- x + 1
		} else {
			quit <- x+1
			return
		}
	}

}

func pong(count int, c1, c2, quit chan int) {
	for {
		time.Sleep(1 * time.Second)
		x := <-c2
		fmt.Println("pong:", x)
		if x < count {
			c1 <- x + 1
		} else {
			quit <- x+1
			return
		}
	}
}
