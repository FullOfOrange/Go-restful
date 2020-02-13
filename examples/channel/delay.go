package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 10)
	
	go func() {
		for {
			time.Sleep(5 * time.Second)
			c1 <- "one"
		}
	}()

	for {
		fmt.Println("start select------------------")
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		}
		fmt.Println("end select-------------------\n\n")
	}
}