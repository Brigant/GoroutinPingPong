package main

import "fmt"

func pingpong(x string,c chan string, quit chan int) {
	for {
		select {
		case c <- x:
			if x == "ping" {
				x = "pong"
			} else {
				x = "ping"
			}

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan string)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	pingpong("ping", c, quit)
}
